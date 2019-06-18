#include "http.h"
#include <windows.h>
#include <cpprest/http_client.h>
#include <cpprest/json.h>
#include <locale.h>
#include <string>
#include <cstring>

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

		postData[L"user_info"][L"name"] = json::value::string(L"this is test");
		postData[L"user_info"][L"sex"] = json::value::string(L"man");
		postData[L"user_info"][L"age"] = json::value::number(20);

		http_client client(L"http://localhost:1213");
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
		return json[L"result"].as_integer();
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