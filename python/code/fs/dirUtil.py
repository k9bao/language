import os

def get_ext(file):
    '''
        返回文件后缀
        如果文件名只包含后缀，比如.gitignore,gradlew，它的后缀是空''
    '''
    return os.path.splitext(file)[-1]

def filter_str(text,in_filter,out_filter):
    '''
    text在out_filter,返回False
    in_filter等于None,返回True
    text在in_filter里边返回True,不在里边返回False
    '''
    if in_filter is not None:
        assert(isinstance(in_filter,list) or isinstance(in_filter,tuple))

    if out_filter is not None:
        assert(isinstance(out_filter,list) or isinstance(out_filter,tuple))

    if out_filter is not None and text in out_filter:
        return False
    if in_filter is not None:
        if text in in_filter:
            return True
        else:
            return False
    else:
        return True

def list_all_files(rootdir,ext=None,non_ext=None):
    '''
    功能：获取目录下的所有文件
    输入：
        rootdir 目录，
        not_ext 排除指定后缀[.json,.txt],优先级高于ext,必须是列表或元祖
        ext     过滤指定后缀，为空不过滤比如[.png,.jpg],只返回.png和.jpg后缀,必须是列表或元祖
    输出：
        文件绝对路径列表
    '''
    _files = []
    list = os.listdir(rootdir)
    for i in range(0,len(list)):
           path = os.path.join(rootdir,list[i])
           if os.path.isdir(path):
              _files.extend(list_all_files(path,ext,non_ext))
           if os.path.isfile(path):
                if filter_str(os.path.splitext(path)[-1],ext,non_ext):
                    _files.append(path)
    return _files

def list_all_ext(rootdir):
    '''
    功能：获取目录下的所有文件夹
    输入：
        rootdir 目录，
    输出：
        所有后缀
    '''
    _ext = set()
    list = os.listdir(rootdir)
    for i in range(0,len(list)):
           path = os.path.join(rootdir,list[i])
           if os.path.isdir(path):
                _ext = _ext.union(list_all_ext(path))
           if os.path.isfile(path):
               _ext.add(os.path.splitext(path)[-1])
    return _ext

if __name__ == "__main__":
    # dir = '.'
    dir = '../../'
    exts = list_all_ext(dir)
    print("后缀总数：{}，类型：{}".format(len(exts),exts))
    files = list_all_files(dir,('.py',))
    print('文件总数：{}'.format(len(files)))
    for f in files:
        print(os.path.basename(f))#打印文件名
    if len(files) > 0:
        f = files[0]
        assert(os.path.exists(f))
        assert(os.path.isfile(f))
        print('\n当前分析文件文件:{}'.format(f))#../../knowledgebao.github.io/_posts\2019-06-28-待整理列表.md
        print("文件目录",os.path.dirname(f))#../../knowledgebao.github.io/_posts
        print("文件名称(含后缀)",os.path.basename(f))#2019-06-28-待整理列表.md
        print("文件名称(不含后缀)",os.path.splitext(os.path.basename(f))[0])#2019-06-28-待整理列表
        print("目录文件名+后缀",os.path.splitext(f))#('../../knowledgebao.github.io/_posts\\2019-06-28-待整理列表', '.md')
        print("后缀(带.)",os.path.splitext(f)[-1])#.md 如果文件名只包含后缀，比如.gitignore,gradlew文件，它的后缀是空''
