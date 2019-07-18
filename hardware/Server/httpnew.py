from http.server import HTTPServer, SimpleHTTPRequestHandler
class MyHandler(SimpleHTTPRequestHandler):
    val_list = [""]
    def __init__(self,request, client_address, server, directory=None):
        super(MyHandler,self).__init__(request, client_address, server)
    def common_response(self):
        self.send_response(200)
        self.send_header('Content-Type', 'text/plain')
        self.end_headers()
        self.wfile.write(b"%s"%(MyHandler.val_list[-1]).encode())

    def do_GET(self):
        self.common_response()

    def do_POST(self):
        content_length = int(self.headers.get('content-length', 0))
        request = self.rfile.read(content_length)  # type: bytes
        print("Request Body: " + request.decode('utf-8'))
        MyHandler.val_list.pop()
        MyHandler.val_list.append(request.decode('utf-8'))
        self.common_response()
        
httpd = HTTPServer(('0.0.0.0', 8000), MyHandler)
print(httpd.server_port)
httpd.serve_forever()