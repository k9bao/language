import re

pattern1 = "(天气)|[^[今明昨前俩三去]{1}(天)[^[前后]{0,1}|(^天)[^[前后啊]{1,1}"
testRe1 = re.compile(pattern1)

if __name__ == "__main__":
    l = ['今天','三天前','俩天后','明天去天津','天后','天啊','天','天前',
        '今天天气怎么样','天气','北京的天气','今天的天怎么样','今天天怎么样','天怎么样',
    ]
    for name in l:
        result = testRe1.findall(name)
        if len(result) > 0:
            print(name,"--------->",result)
