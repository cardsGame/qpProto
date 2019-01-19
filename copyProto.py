import os
import shutil
targetDir = '../clientProto'

def copyProto():
        fileList = os.listdir()
        isExists = os.path.exists(targetDir)
        if not isExists :
                os.makedirs(targetDir)
        for fileName in fileList :
                if fileName.find('.') < 0 and fileName != targetDir:
                        for protoName in os.listdir(fileName) :
                                if protoName.find('.proto') >= 0:
                                        shutil.copy(fileName + '/' + protoName, targetDir)
        
copyProto()
# protoc -I shop --go_out=shop --micro_out=shop shop/shop.proto
# pbjs -t static-module -w commonjs -o ./assets/script/typings/proto.js ./proto/*.proto
# pbjs -t static-module -w commonjs -o ./assets/script/typings/proto.js ./proto/*.proto --no-comments --no-beautify --no-verify --no-convert --no-delimited
