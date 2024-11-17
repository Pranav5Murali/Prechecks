import os
import socket

print("Login name:", os.getenv("USER", "unknown"))
temp = socket.gethostname()
print("Hostname is:", temp)
print("IP details:", socket.gethostbyname(temp))
