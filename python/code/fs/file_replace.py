import re

fileName = "C:/Users/Administrator/Desktop/b2-agents-err.log"
lineList = [
    "history time",
    "A socket operation was attempted to an unreachable host",
    "capture is error Get",
    "Client.Timeout exceeded while awaiting headers",
    "A socket operation was attempted to an unreachable network",
]


def replacePath(file):
    return file.replace('\\', '/')


def delete_line(fileName, lineList):
    matchPatterns = []
    for text in lineList:
        matchPatterns.append(re.compile(text))
    print(fileName, replacePath(fileName))
    with open(replacePath(fileName), mode='r', encoding='utf-8') as f:
        while 1:
            line = f.readline()
            if not line:
                print("Read file over", fileName)
                break
            for p in matchPatterns:
                if p.search(line):
                    break
            else:
                lineList.append(line)

    with open(replacePath(fileName + ".log"), 'w', encoding='UTF-8') as f:
        for i in lineList:
            f.write(i)
        print("Write file over", fileName + ".log")


if __name__ == "__main__":
    delete_line(fileName, lineList)
