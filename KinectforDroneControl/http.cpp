#include "http.h"
#include <windows.h>
#include <iostream>
#include <cpprest/http_client.h>
#include "MyOutputDebugString.h"

using namespace utility;                    // Common utilities like string conversions
using namespace web;                        // Common features like URIs.
using namespace web::http;                  // Common HTTP functionality
using namespace web::http::client;          // HTTP client features
using namespace concurrency::streams;       // Asynchronous streams

#ifdef _DEBUG
#   define MyOutputDebugString( str, ... ) \
      { \
        WCHAR c[256]; \
        swprintf( c, str, __VA_ARGS__ ); \
        OutputDebugString( c ); \
      }
#else
#    define MyOutputDebugString( str, ... ) // ‹óŽÀ‘•
#endif

void prepareHttp() {
	wstring baseUrl = L"http://localhost:1323";
	http_client httpClient(baseUrl);
}

void httpPost()
{
	http_response httpResponse = httpClient.request(methods::POST, L"cancellation", L"This is test").get();

	if (httpResponse.status_code() == status_codes::OK)
	{
		wstring output = httpResponse.extract_utf16string().get();
		const WCHAR* outputWCHAR = output.c_str();
		MyOutputDebugString(outputWCHAR);
	}
}