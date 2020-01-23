import socket
import json
import utils
import threading
 
TCP_IP = '127.0.0.1'
TCP_PORT = 25565
BUFFER_SIZE = 1024
MESSAGE = ""
 
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((TCP_IP, TCP_PORT))

def login(login, password):
    loginJSON = json.dumps({"login":login, "password":password})
    msgJSON = json.dumps({"type":"login", "data": loginJSON})
    s.send(msgJSON.encode('utf-8'))

def register(login, password):
    loginJSON = json.dumps({"login":login, "password":password})
    msgJSON = json.dumps({"type":"register", "data": loginJSON})
    s.send(msgJSON.encode('utf-8'))

def send_all(msg):
    sendAllJSON = json.dumps({"message":msg})
    msgJSON = json.dumps({"type":"send_all", "data": sendAllJSON})
    s.send(msgJSON.encode('utf-8'))

def receive():
    try:
        netData = utils.receiveJSON(s)
        
        # print(''.join(netData))
        msg = json.loads(''.join(netData))
        # print(msg)
        login = json.loads(msg['data'])
        print(login)
    except :
        print("Error!")

def receiver():
    while True:
        receive()

x = threading.Thread(target=receiver)
x.start()

register("alex", "123")

login("alex", "123")

login("alex", "123")

register("alex2", "123")

login("alex2", "123")


send_all("Take it boy")

while True:
    pass

s.close()