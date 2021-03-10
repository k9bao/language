import numpy as np
import cv2
import os

def list_all_files(rootdir,extern=[]):
    '''
    功能：获取目录下的所有文件
    输入：
        rootdir 目录，
        extern 后缀过滤列表，为空不过滤比如[.png,.jpg]
    输出：
        文件绝对路径列表
    '''
    _files = []
    list = os.listdir(rootdir)
    for i in range(0,len(list)):
           path = os.path.join(rootdir,list[i])
           if os.path.isdir(path):
              _files.extend(list_all_files(path))
           if os.path.isfile(path):
                if len(extern) > 0:
                   if os.path.splitext(path)[-1] in extern:
                    _files.append(path)
                else:
                    _files.append(path)
              
    return _files


def cv_imread(filePath):
    '''
    读取图像，解决imread不能读取中文路径的问题
    '''
    cv_img=cv2.imdecode(np.fromfile(filePath,dtype=np.uint8),-1)
    ## imdecode读取的是rgb，如果后续需要opencv处理的话，需要转换成bgr，转换后图片颜色会变化
    ##cv_img=cv2.cvtColor(cv_img,cv2.COLOR_RGB2BGR)
    return cv_img

def RotateClockWise90(img):
    '''
    将图片顺时针旋转90度
    '''
    trans_img = cv2.transpose( img )#转置阵图片(镜像图片)
    # new_img = cv2.flip(trans_img, 0)#实现沿X轴方向的镜像图片,将图片顺时针旋转90度
    new_img = cv2.flip(trans_img, 1)#实现沿Y轴方向的镜像图片，,将图片逆时针旋转90度
    return new_img

if __name__ == "__main__":
    extern = ['.png']
    files = list_all_files("C:\\Users\\Administrator\\Desktop\\pdf_to_image.winodws.x86_64\\small",extern)
    for f in files:
        print(f)
        img=cv_imread(f)
        img90=RotateClockWise90(img)
        cv2.imencode('.png', img90)[1].tofile(f)
 