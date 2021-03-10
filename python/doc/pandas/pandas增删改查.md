
# 1. Pandas基本结构增删改查

- [1. Pandas基本结构增删改查](#1-pandas基本结构增删改查)
  - [1.1. 基本结构 Series 和 DataFrame](#11-基本结构-series-和-dataframe)
  - [1.2. Series的增删改查](#12-series的增删改查)
    - [1.2.1. 索引操作](#121-索引操作)
    - [1.2.2. 数据操作](#122-数据操作)
  - [1.3. DataFrame的增删改查](#13-dataframe的增删改查)
    - [1.3.1. 索引操作](#131-索引操作)
    - [1.3.2. 数据操作](#132-数据操作)
    - [1.3.3. 遍历](#133-遍历)

## 1.1. 基本结构 Series 和 DataFrame


```python
# pandas基本结构，Series和DataFrame
import numpy as np
import pandas as pd
from IPython.core.interactiveshell import InteractiveShell

InteractiveShell.ast_node_interactivity = 'all' #使得独占一行的所有变量或者语句都自动打印显示，默认只显示最后一行
        
series = pd.Series([20,30,40])
series
```




    0    20
    1    30
    2    40
    dtype: int64




```python
stock_data = pd.DataFrame([[1,2,3,4],[5,6,7,8],[9,10,11,12]])
stock_data
```




<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>0</th>
      <th>1</th>
      <th>2</th>
      <th>3</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>1</td>
      <td>2</td>
      <td>3</td>
      <td>4</td>
    </tr>
    <tr>
      <th>1</th>
      <td>5</td>
      <td>6</td>
      <td>7</td>
      <td>8</td>
    </tr>
    <tr>
      <th>2</th>
      <td>9</td>
      <td>10</td>
      <td>11</td>
      <td>12</td>
    </tr>
  </tbody>
</table>
</div>



## 1.2. Series的增删改查

### 1.2.1. 索引操作


```python
#修改索引
series.index=['a','b','c']
# series.reindex(index=['a','b','c'])
series
```




    a    20
    b    30
    c    40
    dtype: int64




```python
#查看索引
series.index
```




    Index(['a', 'b', 'c'], dtype='object')



### 1.2.2. 数据操作


```python
# 查看
series['a']
```




    20




```python
# 修改
series['a']=100
series['a']
```




    100




```python
# 增加
series_tmp = pd.Series(index=['e','f'],data=[11,12])
series1 = series.append(series_tmp)#返回值是处理后对象，原对象不变
series1
series
```




    a    100
    b     30
    c     40
    e     11
    f     12
    dtype: int64






    a    100
    b     30
    c     40
    dtype: int64




```python
# 删除
series2 = series.drop("c")#返回值是处理后对象，原对象不变
series2
series
```




    a    100
    b     30
    dtype: int64






    a    100
    b     30
    c     40
    dtype: int64



## 1.3. DataFrame的增删改查

### 1.3.1. 索引操作

一般常用的有两个方法：
1、使用DataFrame.index = [newName]，DataFrame.columns = [newName]，这两种方法可以轻松实现。
2、使用rename方法（推荐）：
DataFrame.rename（mapper = None，index = None，columns = None，axis = None，copy = True，inplace = False，level = None ）
参数介绍：
mapper，index，columns：可以任选其一使用，可以是将index和columns结合使用。index和column直接传入mapper或者字典的形式。
axis：int或str，与mapper配合使用。可以是轴名称（‘index’，‘columns’）或数字（0,1）。默认为’index’。
copy：boolean，默认为True，是否复制基础数据。
inplace：布尔值，默认为False，是否返回新的DataFrame。如果为True，则忽略复制值。


```python
df1 = pd.DataFrame(np.arange(9).reshape(3, 3), index = ['bj', 'sh', 'gz'], columns=['a', 'b', 'c'])
df1
df1.index = pd.Series(['beijing', 'shanghai', 'guangzhou'])# 直接修改
df1
df1.rename(index=str.upper, columns=str.lower,inplace=False) # 返回新DataFrame
df1.rename(index={'beijing':'bj'}, columns = {'a':'aa'},inplace=False) # 改某一个标题，返回新DataFrame
```




<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>a</th>
      <th>b</th>
      <th>c</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>bj</th>
      <td>0</td>
      <td>1</td>
      <td>2</td>
    </tr>
    <tr>
      <th>sh</th>
      <td>3</td>
      <td>4</td>
      <td>5</td>
    </tr>
    <tr>
      <th>gz</th>
      <td>6</td>
      <td>7</td>
      <td>8</td>
    </tr>
  </tbody>
</table>
</div>






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>a</th>
      <th>b</th>
      <th>c</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>beijing</th>
      <td>0</td>
      <td>1</td>
      <td>2</td>
    </tr>
    <tr>
      <th>shanghai</th>
      <td>3</td>
      <td>4</td>
      <td>5</td>
    </tr>
    <tr>
      <th>guangzhou</th>
      <td>6</td>
      <td>7</td>
      <td>8</td>
    </tr>
  </tbody>
</table>
</div>






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>a</th>
      <th>b</th>
      <th>c</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>BEIJING</th>
      <td>0</td>
      <td>1</td>
      <td>2</td>
    </tr>
    <tr>
      <th>SHANGHAI</th>
      <td>3</td>
      <td>4</td>
      <td>5</td>
    </tr>
    <tr>
      <th>GUANGZHOU</th>
      <td>6</td>
      <td>7</td>
      <td>8</td>
    </tr>
  </tbody>
</table>
</div>






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>aa</th>
      <th>b</th>
      <th>c</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>bj</th>
      <td>0</td>
      <td>1</td>
      <td>2</td>
    </tr>
    <tr>
      <th>shanghai</th>
      <td>3</td>
      <td>4</td>
      <td>5</td>
    </tr>
    <tr>
      <th>guangzhou</th>
      <td>6</td>
      <td>7</td>
      <td>8</td>
    </tr>
  </tbody>
</table>
</div>




```python
def test_map(x):
    return x+'_ABC'
df1.index.map(test_map)
df1.rename(index=test_map,columns=test_map)
```




    Index(['beijing_ABC', 'shanghai_ABC', 'guangzhou_ABC'], dtype='object')






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>a_ABC</th>
      <th>b_ABC</th>
      <th>c_ABC</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>beijing_ABC</th>
      <td>0</td>
      <td>1</td>
      <td>2</td>
    </tr>
    <tr>
      <th>shanghai_ABC</th>
      <td>3</td>
      <td>4</td>
      <td>5</td>
    </tr>
    <tr>
      <th>guangzhou_ABC</th>
      <td>6</td>
      <td>7</td>
      <td>8</td>
    </tr>
  </tbody>
</table>
</div>



### 1.3.2. 数据操作

    参数inplace默认为False,只能在生成的新数据块中实现编辑效果。当inplace=True时执行内部编辑，不返回任何值，原数据发生改变。


```python
#增加列
df = pd.DataFrame(data = [['lisa','f',22],['joy','f',22],['tom','m','21']],index = ['a','b','c'],columns = ['name','sex','age'])
df
df.insert(0,'city',['ny','zz','xy']) #在第0列，加上column名称为city列。
df['job'] = ['student','AI','teacher'] #默认在df最后一列加上column名称为job的列
df.loc[:,'salary'] = ['1k','2k','3k'] #在df最后一列加上column名称为salary列
df
```




<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>name</th>
      <th>sex</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>lisa</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>b</th>
      <td>joy</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>c</th>
      <td>tom</td>
      <td>m</td>
      <td>21</td>
    </tr>
  </tbody>
</table>
</div>






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>city</th>
      <th>name</th>
      <th>sex</th>
      <th>age</th>
      <th>job</th>
      <th>salary</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>ny</td>
      <td>lisa</td>
      <td>f</td>
      <td>22</td>
      <td>student</td>
      <td>1k</td>
    </tr>
    <tr>
      <th>b</th>
      <td>zz</td>
      <td>joy</td>
      <td>f</td>
      <td>22</td>
      <td>AI</td>
      <td>2k</td>
    </tr>
    <tr>
      <th>c</th>
      <td>xy</td>
      <td>tom</td>
      <td>m</td>
      <td>21</td>
      <td>teacher</td>
      <td>3k</td>
    </tr>
  </tbody>
</table>
</div>




```python
## 增加行
df_insert = pd.DataFrame({'name':['mason','mario'],'sex':['m','f'],'age':[21,22]},index = [4,5])
df_insert   
df = pd.DataFrame(data = [['lisa','f',22],['joy','f',22],['tom','m','21']],index = ['a','b','c'],columns = ['name','sex','age'])
#返回添加后的值，并不会修改df的值。
# ignore_index默认为False，意思是不忽略index值，即生成的新的ndf的index采用df_insert中的index值。
# 若为True，则新的ndf的index值不使用df_insert中的index值，而是自己默认生成。默认会排序
df.append(df_insert,ignore_index = False,sort=False) 
df
```




<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>name</th>
      <th>sex</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>4</th>
      <td>mason</td>
      <td>m</td>
      <td>21</td>
    </tr>
    <tr>
      <th>5</th>
      <td>mario</td>
      <td>f</td>
      <td>22</td>
    </tr>
  </tbody>
</table>
</div>






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>name</th>
      <th>sex</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>lisa</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>b</th>
      <td>joy</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>c</th>
      <td>tom</td>
      <td>m</td>
      <td>21</td>
    </tr>
    <tr>
      <th>4</th>
      <td>mason</td>
      <td>m</td>
      <td>21</td>
    </tr>
    <tr>
      <th>5</th>
      <td>mario</td>
      <td>f</td>
      <td>22</td>
    </tr>
  </tbody>
</table>
</div>






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>name</th>
      <th>sex</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>lisa</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>b</th>
      <td>joy</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>c</th>
      <td>tom</td>
      <td>m</td>
      <td>21</td>
    </tr>
  </tbody>
</table>
</div>




```python
# 查改 loc使用index或columns查找，iloc使用对应的数字索引查找。使用loc和iloc查找后可以直接修改
df = pd.DataFrame(data = [['lisa','f',22],['joy','f',22],['tom','m','21']],index = ['a','b','c'],columns = ['name','sex','age'])
df
df.loc['a','name']                      # 选取某个值
df.iloc[1]                              # 选取某一行
df.loc[:,'name']==df.name               # 选取某一列
df.iloc[0:2, [0,2]]                     #选取多行多列，iloc参数[0:2]不包含2
df.loc['a':'b',['name','age']]          #选取多行多列，loc参数['a':'b']包含'b'
df.loc[df['sex']=='m','name']           #选取gender列是m，name列的数据
df.loc[df['sex']=='m',['name','age']]   #选取gender列是m，name和age列的数据
```




<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>name</th>
      <th>sex</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>lisa</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>b</th>
      <td>joy</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>c</th>
      <td>tom</td>
      <td>m</td>
      <td>21</td>
    </tr>
  </tbody>
</table>
</div>






    'lisa'






    name    joy
    sex       f
    age      22
    Name: b, dtype: object






    a    True
    b    True
    c    True
    Name: name, dtype: bool






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>name</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>lisa</td>
      <td>22</td>
    </tr>
    <tr>
      <th>b</th>
      <td>joy</td>
      <td>22</td>
    </tr>
  </tbody>
</table>
</div>






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>name</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>lisa</td>
      <td>22</td>
    </tr>
    <tr>
      <th>b</th>
      <td>joy</td>
      <td>22</td>
    </tr>
  </tbody>
</table>
</div>






    c    tom
    Name: name, dtype: object






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>name</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>c</th>
      <td>tom</td>
      <td>21</td>
    </tr>
  </tbody>
</table>
</div>




```python
df = pd.DataFrame(data = [['lisa','f',22],['joy','f',22],['tom','m','21']],index = ['a','b','c'],columns = ['name','sex','age'])
df
df.drop(['a','c'],axis = 0,inplace = False)#删除index值为a和c的两行，
```




<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>name</th>
      <th>sex</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>lisa</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>b</th>
      <td>joy</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>c</th>
      <td>tom</td>
      <td>m</td>
      <td>21</td>
    </tr>
  </tbody>
</table>
</div>






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>name</th>
      <th>sex</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>b</th>
      <td>joy</td>
      <td>f</td>
      <td>22</td>
    </tr>
  </tbody>
</table>
</div>




```python
df = pd.DataFrame(data = [['lisa','f',22],['joy','f',22],['tom','m','21']],index = ['a','b','c'],columns = ['name','sex','age'])
df
df.drop(['name'],axis = 1,inplace = False)  #删除name列，新返回的DataFrame无name列
df
del df['name']  #删除name列,相当于inplace = True，这一列无输出
df
df.pop('age')  #删除age列，操作后，df丢掉了age列,age列作为返回值返回
df
```




<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>name</th>
      <th>sex</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>lisa</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>b</th>
      <td>joy</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>c</th>
      <td>tom</td>
      <td>m</td>
      <td>21</td>
    </tr>
  </tbody>
</table>
</div>






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>sex</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>b</th>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>c</th>
      <td>m</td>
      <td>21</td>
    </tr>
  </tbody>
</table>
</div>






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>name</th>
      <th>sex</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>lisa</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>b</th>
      <td>joy</td>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>c</th>
      <td>tom</td>
      <td>m</td>
      <td>21</td>
    </tr>
  </tbody>
</table>
</div>






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>sex</th>
      <th>age</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>b</th>
      <td>f</td>
      <td>22</td>
    </tr>
    <tr>
      <th>c</th>
      <td>m</td>
      <td>21</td>
    </tr>
  </tbody>
</table>
</div>






    a    22
    b    22
    c    21
    Name: age, dtype: object






<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>sex</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>a</th>
      <td>f</td>
    </tr>
    <tr>
      <th>b</th>
      <td>f</td>
    </tr>
    <tr>
      <th>c</th>
      <td>m</td>
    </tr>
  </tbody>
</table>
</div>



### 1.3.3. 遍历


```python
inp = [{'c1':10, 'c2':100}, {'c1':11, 'c2':110}, {'c1':12, 'c2':123}]
df = pd.DataFrame(inp)
df
#按照行遍历
for index, row in df.iterrows():
    print(index,row['c2']) # 输出每行的索引值
```




<div>
<style scoped>
    .dataframe tbody tr th:only-of-type {
        vertical-align: middle;
    }

    .dataframe tbody tr th {
        vertical-align: top;
    }

    .dataframe thead th {
        text-align: right;
    }
</style>
<table border="1" class="dataframe">
  <thead>
    <tr style="text-align: right;">
      <th></th>
      <th>c1</th>
      <th>c2</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <th>0</th>
      <td>10</td>
      <td>100</td>
    </tr>
    <tr>
      <th>1</th>
      <td>11</td>
      <td>110</td>
    </tr>
    <tr>
      <th>2</th>
      <td>12</td>
      <td>123</td>
    </tr>
  </tbody>
</table>
</div>



    0 100
    1 110
    2 123
    


```python
#按照列遍历
for index, row in df.iteritems():
    print(index,row[0], row[2]) # 输出列名
```

    c1 10 12
    c2 100 123
    


```python

```
