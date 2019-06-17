#pragma once

#include <string>
using namespace std;

#include "http.h"
#include <windows.h>
#include <iostream>
#include <cpprest/http_client.h>

using namespace utility;                    // Common utilities like string conversions
using namespace web;                        // Common features like URIs.
using namespace web::http;                  // Common HTTP functionality
using namespace web::http::client;          // HTTP client features
using namespace concurrency::streams;       // Asynchronous streams


 extern void httpPost();

 extern wstring baseUrl;
 extern http_client httpClient;
