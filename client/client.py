import socket
import json
 
TCP_IP = '127.0.0.1'
TCP_PORT = 25565
BUFFER_SIZE = 1024
MESSAGE = ""
 
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect((TCP_IP, TCP_PORT))

def send(login, password):
    loginJSON = json.dumps({"Login":login, "Password":password})
    msgJSON = json.dumps({"MessageType":"login", "Message": loginJSON})
    s.send(msgJSON.encode('utf-8'))

def receive():
    try:
        netData = []
        balance = 0
        b = s.recv(1).decode('utf-8')
        if b != r'{':
            return

        netData.append(b)
        balance += 1
        while balance > 0:
            b = s.recv(1).decode('utf-8')
            
            netData.append(b)

            if b == r'{':
                balance+=1
            elif b == r'}':
                balance-=1
        print(''.join(netData))
        msg = json.loads(''.join(netData))
        print(msg)
        login = json.loads(msg['Message'])
        print(login)
    except :
        print("Error!")

send("123", "123")
print(0)
receive()
print(0)
send("pl123", "pl123")
print(0)
receive()
print(0)

s.close()