#include "http.h"
#include <windows.h>
#include <cpprest/http_client.h>
#include <cpprest/filestream.h>
#include <iostream>
//#include <string>
//#include <cstring>
constexpr auto SCALING_NUMBER = 1000000;

using namespace utility;
using namespace web;                        // Common features like URIs.
using namespace web::http;                  // Common HTTP functionality
using namespace web::http::client;          // HTTP client features


#ifdef _DEBUG
#   define MyOutputDebugString( str, ... ) \
      { \
        WCHAR c[256]; \
        swprintf( c, str, __VA_ARGS__ ); \
        OutputDebugString( c ); \
      }
#else
#    define MyOutputDebugString( str, ... ) // �����
#endif


pplx::task<int> Post(int rx, int ry, int rz, int rshape,
					 int lx, int ly, int lz, int lshape)
{
	return pplx::create_task([&]
	{
		json::value postData;

		postData[L"right"][L"x"] = json::value::number(rx);
		postData[L"right"][L"y"] = json::value::number(ry);
		postData[L"right"][L"z"] = json::value::number(rz);
		postData[L"right"][L"shape"] = json::value::number(rshape);
		postData[L"left"][L"x"] = json::value::number(lx);
		postData[L"left"][L"y"] = json::value::number(ly);
		postData[L"left"][L"z"] = json::value::number(lz);
		postData[L"left"][L"shape"] = json::value::number(lshape);

		http_client client(L"http://localhost:1323");
		return client.request(methods::POST, L"", postData.serialize(), L"application/json");
	}).then([](http_response response)
	{
		if (response.status_code() == status_codes::OK)
		{
			return 200;
		}
		else {
			return 500;
		}
		//if (response.status_code() == status_codes::OK)
		//{
		//	//auto body = response.extract_string();
		//	//std::wcout << body.get().c_str() << std::endl;
		//	//std::cout << response.extract_json() << std::endl;
		//	return response.extract_json();
		//}
	//}).then([](json::value json)
	//{
	//	
	//	// ���U���g�R�[�h��Ԃ�
	//	int resultCode = json[L"time"].as_integer();
	//	MyOutputDebugString(L"time = %d\n", resultCode);
	//	return resultCode;
	//	
	});
}

void httpPost(double rx, double ry, double rz, int rshape,
			  double lx, double ly, double lz, int lshape)
{
	try
	{
		int _rx = rx * SCALING_NUMBER;
		int _ry = ry * SCALING_NUMBER;
		int _rz = rz * SCALING_NUMBER;
		int _lx = lx * SCALING_NUMBER;
		int _ly = ly * SCALING_NUMBER;
		int _lz = lz * SCALING_NUMBER;
		auto result = Post(_rx, _ry, _rz, rshape, _lx, _ly, _lz, lshape).wait();
		MyOutputDebugString(L"Result=%d\n", result);
	}
	catch (const std::exception& e)
	{
		OutputDebugString(L"Error");
		WCHAR ws[100];

		mbstowcs(ws, e.what(), 100);
		OutputDebugString(ws);
		OutputDebugString(L"\n");
	}
}