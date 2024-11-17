import os
import socket

print("Login name:",os.getlogin())
temp=socket.gethostname()
print("hostname is:",temp)
print("IP_details",socket.gethostbyname(temp))

