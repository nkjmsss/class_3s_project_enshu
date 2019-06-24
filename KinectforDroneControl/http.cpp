#include "http.h"
#include <windows.h>
#include <cpprest/http_client.h>
#include <cpprest/filestream.h>
#include <iostream>
//#include <string>
//#include <cstring>

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
#    define MyOutputDebugString( str, ... ) // 空実装
#endif


pplx::task<int> Post()
{
	return pplx::create_task([]
	{
		json::value postData;

		postData[L"time"] = json::value::number(1653);
		postData[L"x"] = json::value::number(20);
		postData[L"y"] = json::value::number(40);
		postData[L"z"] = json::value::number(60);
		postData[L"shape"] = json::value::number(2);

		http_client client(L"http://localhost:1323");
		return client.request(methods::POST, L"", postData.serialize(), L"application/json");
	}).then([](http_response response)
	{
		if (response.status_code() == status_codes::OK)
		{
			//auto body = response.extract_string();
			//std::wcout << body.get().c_str() << std::endl;
			//std::cout << response.extract_json() << std::endl;
			return response.extract_json();
		}
	}).then([](json::value json)
	{
		
		// リザルトコードを返す
		int resultCode = json[L"time"].as_integer();
		MyOutputDebugString(L"time = %d\n", resultCode);
		return resultCode;
		
	});
}

void httpPost()
{
	try
	{
		auto result = Post().wait();
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