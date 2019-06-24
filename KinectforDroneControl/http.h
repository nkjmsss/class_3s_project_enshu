#pragma once

#include <windows.h>
#include <iostream>
#include <cpprest/http_client.h>


 extern void httpPost(double rx, double ry, double rz, int rshape,
					  double lx, double ly, double lz, int lshape);

 extern web::http::client::http_client httpClient;
