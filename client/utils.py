import socket

def receiveJSON(sock):
    netData = []
    balance = 0
    inString = False
    passByte = False
    b = sock.recv(1).decode('utf-8')
    if b != r'{':
        return

    netData.append(b)
    balance += 1
    while balance > 0:
        b = sock.recv(1).decode('utf-8')
        
        netData.append(b)

        if passByte:
            passByte = False
        else:
            if b == '\\':
                passByte = True
            if b == '"':
                inString = not inString
            if not inString:
                if b == '{':
                    balance += 1
                if b == '}':
                    balance -= 1

    return netData