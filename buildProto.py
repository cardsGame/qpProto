import os
targetFolder = './assets/script/typings/'


def run(coroutine):
    try:
        coroutine.send(None)
    except StopIteration as e:
        return e.value


def removeOldTypings():
    for fileName in os.listdir(targetFolder):
        path_file = os.path.join(targetFolder, fileName)
        os.path.isfile(path_file)
        os.remove(path_file)


def genProto():
    srcStr = r'\src'
    splitStr = ';'
    if os.name != 'nt':
        os.system('source ~/.bash_profile')
        splitStr = ':'
        srcStr = r'/src'
    print('操作系统:', os.name)
    fileList = os.listdir()
    # 关闭goModule先
    os.environ['GO111MODULE'] = 'off'
    # 通过goget拉取项目路径到GOPATH
    os.system('go get -u -x github.com/cardsGame/qpProto')
    goPath = os.getenv('GOPATH') or '/Users/pengju/go/src'
    print("goPath:", goPath)
    dIndex = goPath.find(splitStr)
    print("dIndex:", dIndex)
    if (dIndex > 0):
        goPath = goPath[:dIndex]
    print("newGoPath:", goPath)
    goPath = goPath + srcStr
    folderList = []
    for i in range(0, len(fileList)):
        fileName = fileList[i]
        dotIndex = fileName.find('.')
        if (dotIndex < 0):
            folderList.append(fileName)
    print("folderList:", folderList)
    for folderName in folderList:
        cmd = 'protoc --proto_path="' + goPath + \
            '" --proto_path=. --go_out=plugins=grpc:. --micro_out=plugins=grpc:. ' + \
            folderName + '/' + folderName + '.proto'
        os.system(cmd)


genProto()
# protoc -I shop --go_out=shop --micro_out=shop shop/shop.proto
# pbjs -t static-module -w commonjs -o ./assets/script/typings/proto.js ./proto/*.proto
# pbjs -t static-module -w commonjs -o ./assets/script/typings/proto.js ./proto/*.proto --no-comments --no-beautify --no-verify --no-convert --no-delimited
