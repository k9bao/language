# 1. datetime

- [1. datetime](#1-datetime)
  - [1.1. 简介](#11-简介)
  - [1.2. datetime.timedelta](#12-datetimetimedelta)
  - [1.3. datetime.datetime](#13-datetimedatetime)
    - [1.3.1. 类函数](#131-类函数)
    - [1.3.2. 类属性](#132-类属性)
    - [1.3.3. 实例属性](#133-实例属性)
    - [1.3.4. 支持的运算](#134-支持的运算)
    - [1.3.5. 实例方法](#135-实例方法)
  - [1.4. datetime.time](#14-datetimetime)
    - [1.4.1. 类方法](#141-类方法)
    - [1.4.2. 类属性](#142-类属性)
    - [1.4.3. 实例属性](#143-实例属性)
    - [1.4.4. 实例方法](#144-实例方法)
  - [1.5. datetime.date](#15-datetimedate)
    - [1.5.1. 类方法](#151-类方法)
    - [1.5.2. 实例方法](#152-实例方法)
    - [1.5.3. 类属性](#153-类属性)
    - [1.5.4. 实例属性](#154-实例属性)
    - [1.5.5. 支持的运算](#155-支持的运算)
  - [1.6. datetime.tzinfo](#16-datetimetzinfo)
    - [1.6.1. datetime.timezone](#161-datetimetimezone)
  - [1.7. strftime() 和 strptime() 的行为](#17-strftime-和-strptime-的行为)
  - [1.8. 参考资料](#18-参考资料)

## 1.1. 简介

Unix时间戳(Unix timestamp)，或称Unix时间(Unix time)、POSIX时间(POSIX time)，是一种时间表示方式，定义为从格林威治时间1970年01月01日00时00分00秒起至现在的总秒数。

类继承结构

```txt
object
    timedelta
    tzinfo
        timezone
    time
    date
        datetime
```

## 1.2. datetime.timedelta

&emsp;&emsp;表示两个 date 对象或者 time 对象,或者 datetime 对象之间的时间间隔，精确到微秒。

## 1.3. datetime.datetime

&emsp;&emsp;日期和时间的结合。属性：year, month, day, hour, minute, second, microsecond, and tzinfo.datetime 对象是包含来自 date 对象和 time 对象的所有信息的单一对象。与 date 对象一样，datetime 假定当前的格列高利历向前后两个方向无限延伸；与 time 对象一样，datetime 假定每一天恰好有 3600*24 秒。

### 1.3.1. 类函数

- `class datetime.datetime(year, month, day, hour=0, minute=0, second=0, microsecond=0, tzinfo=None, *, fold=0)`
  - year, month 和 day 参数是必须的。 tzinfo 可以是 None 或者是一个 tzinfo 子类的实例。 其余的参数必须是在下面范围内的整数：
  - fold in [0, 1].如果参数超出范围，则抛出 ValueError 异常。3.6 新版功能: 增加了 fold 参数。

- `classmethod datetime.today()` 使用 `now()` 代替
- `classmethod datetime.utcnow()` 使用 `datetime.now(timezone.utc)` 代替
- `classmethod datetime.now(tz=None)`
  - 返回表示当前地方时的 date 和 time 对象。
  - 如果可选参数 tz 为 None 或未指定，这就类似于 today()，但该方法会在可能的情况下提供比通过 time.time() 时间戳所获时间值更高的精度（例如，在提供了 C gettimeofday() 函数的平台上就可以做到这一点）。
  - 如果 tz 不为 None，它必须是 tzinfo 子类的一个实例，并且当前日期和时间将被转换到 tz 时区。
  - 此函数可以替代 today() 和 utcnow()。

- `classmethod datetime.utcfromtimestamp(timestamp)` 使用 `datetime.fromtimestamp(timestamp, tz=timezone.utc)` 代替
  - 返回对应于 POSIX 时间戳的 UTC datetime，其中 tzinfo 值为 None。（结果为简单型对象。）
- `classmethod datetime.fromtimestamp(timestamp, tz=None)`
  - 返回对应于 POSIX 时间戳例如 time.time() 的返回值的本地日期和时间。 如果可选参数 tz 为 None 或未指定，时间戳会被转换为所在平台的本地日期和时间，返回的 datetime 对象将为感知型。
  - 如果 tz 不为 None，它必须是 tzinfo 子类的一个实例，并且时间戳将被转换到 tz 指定的时区。
  - fromtimestamp() 可能会引发 OverflowError，如果时间戳数值超出所在平台 C localtime() 或 gmtime() 函数的支持范围的话，并会在 localtime() 或 gmtime() 报错时引发 OSError。 通常该数值会被限制在 1970 年至 2038 年之间。 请注意在时间戳概念包含闰秒的非 POSIX 系统上，闰秒会被 fromtimestamp() 所忽略，结果可能导致两个相差一秒的时间戳产生相同的 datetime 对象。 相比 utcfromtimestamp() 更推荐使用此方法。
  - 在 3.3 版更改: 引发 OverflowError 而不是 ValueError，如果时间戳数值超出所在平台 C localtime() 或 gmtime() 函数的支持范围的话。 并会在 localtime() 或 gmtime() 出错时引发 OSError 而不是 ValueError。
  - 在 3.6 版更改: fromtimestamp() 可能返回 fold 值设为 1 的实例。

- `classmethod datetime.fromordinal(ordinal)` 参考 `classmethod date.fromordinal(ordinal)`

- `classmethod datetime.combine(date, time, tzinfo=self.tzinfo)`
  - 返回一个新的 datetime 对象，对象的日期部分等于给定的 date 对象的值，而其时间部分等于给定的 time 对象的值。 如果提供了 tzinfo 参数，其值会被用来设置结果的 tzinfo 属性，否则将使用 time 参数的 tzinfo 属性。
  - 在 3.6 版更改: 增加了 tzinfo 参数。

- `classmethod datetime.fromisocalendar(year, week, day)` 参考 `classmethod date.fromisocalendar(year, week, day)`，其他取默认值

- `classmethod datetime.fromisoformat(date_string)`
  - 返回一个对应于 date.isoformat() 和 datetime.isoformat() 所提供的某一种 date_string 的 datetime 对象。
  - 3.7 新版功能.

- `classmethod datetime.strptime(date_string, format)`
  - 返回一个对应于 date_string，根据 format 进行解析得到的 datetime 对象。
  - `datetime(*(time.strptime(date_string, format)[0:6]))`
  - 如果 date_string 和 format 无法被 time.strptime() 解析或它返回一个不是时间元组的值，则将引发 ValueError。 要获取格式化指令的完整列表，请参阅 strftime() 和 strptime() 的行为。

### 1.3.2. 类属性

- `datetime.min`
  - 最早的可表示 datetime，datetime(MINYEAR, 1, 1, tzinfo=None)。

- `datetime.max`
  - 最晚的可表示 datetime，datetime(MAXYEAR, 12, 31, 23, 59, 59, 999999, tzinfo=None)。

- `datetime.resolution`
  - 两个不相等的 datetime 对象之间可能的最小间隔，timedelta(microseconds=1)。

### 1.3.3. 实例属性

`datetime.year`
在 MINYEAR 和 MAXYEAR 之间，包含边界。

`datetime.month`
1 至 12（含）

`datetime.day`
返回1到指定年月的天数间的数字。

`datetime.hour`
取值范围是 range(24)。

`datetime.minute`
取值范围是 range(60)。

`datetime.second`
取值范围是 range(60)。

`datetime.microsecond`
取值范围是 range(1000000)。

`datetime.tzinfo`
作为 tzinfo 参数被传给 datetime 构造器的对象，如果没有传入值则为 None。

`datetime.fold`
取值范围是 [0, 1]。 用于在重复的时间段中消除边界时间歧义。 （当夏令时结束时回拨时钟或由于政治原因导致当明时区的 UTC 时差减少就会出现重复的时间段。） 取值 0 (1) 表示两个时刻早于（晚于）所代表的同一边界时间。

### 1.3.4. 支持的运算

- `datetime2 = datetime1 + timedelta`
  - datetime2 是从 datetime1 去掉了一段 timedelta 的结果，如果 timedelta.days > 0 则是在时间线上前进，如果 timedelta.days < 0 则是在时间线上后退。 该结果具有与输入的 datetime 相同的 tzinfo 属性，并且操作完成后 datetime2 - datetime1 == timedelta。 如果 datetime2.year 将要小于 MINYEAR 或大于 MAXYEAR 则会引发 OverflowError。 请注意即使输入的是一个感知型对象，该方法也不会进行时区调整。
- `datetime2 = datetime1 - timedelta`
  - 计算 datetime2 使得 datetime2 + timedelta == datetime1。 与相加操作一样，结果具有与输入的 datetime 相同的 tzinfo 属性，即使输入的是一个感知型对象，该方法也不会进行时区调整。
- `timedelta = datetime1 - datetime2`
  - 从一个 datetime 减去一个 datetime 仅对两个操作数均为简单型或均为感知型时有定义。 如果一个是感知型而另一个是简单型，则会引发 TypeError。
  - 如果两个操作数都是简单型，或都是感知型并且具有相同的 tzinfo 属性，则 tzinfo 属性会被忽略，并且结果会是一个使得 datetime2 + t == datetime1 的 timedelta 对象 t。 在此情况下不会进行时区调整。
  - 如果两个操作数都是感知型且具有不同的 tzinfo 属性，a-b 操作的效果就如同 a 和 b 首先被转换为简单型 UTC 日期时间。 结果将是 (a.replace(tzinfo=None) - a.utcoffset()) - (b.replace(tzinfo=None) - b.utcoffset())，除非具体实现绝对不溢出。

- `datetime1 < datetime2`
  - 比较 datetime 与 datetime
  - 当 datetime1 的时间在 datetime2 之前则认为 datetime1 小于 datetime2。
  - 如果比较的一方是简单型而另一方是感知型，则如果尝试进行顺序比较将引发 TypeError。 对于相等比较，简单型实例将永远不等于感知型实例。
  - 如果两个比较方都是感知型，且具有相同的 tzinfo 属性，则相同的 tzinfo 属性会被忽略并对基本日期时间值进行比较。 如果两个比较方都是感知型且具有不同的 tzinfo 属性，则两个比较方将首先通过减去它们的 UTC 差值（使用 self.utcoffset() 获取）来进行调整。
  - 在 3.3 版更改: 感知型和简单型 datetime 实例之间的相等比较不会引发 TypeError。

  注解 为了防止比较操作回退为默认的对象地址比较方式，datetime 比较通常会引发 TypeError，如果比较目标不同样为 datetime 对象的话。 不过也可能会返回 NotImplemented，如果比较目标具有 timetuple() 属性的话。 这个钩子给予其他种类的日期对象实现混合类型比较的机会。 如果未实现，则当 datetime 对象与不同类型比较时将会引发 TypeError，除非是 == 或 != 比较。 后两种情况将分别返回 False 或 True。

### 1.3.5. 实例方法

`datetime.date()`
返回具有同样 year, month 和 day 值的 date 对象。

`datetime.time()`
返回具有同样 hour, minute, second, microsecond 和 fold 值的 time 对象。 tzinfo 值为 None。 另请参见 timetz() 方法。

在 3.6 版更改: fold 值会被复制给返回的 time 对象。

`datetime.timetz()`
返回具有同样 hour, minute, second, microsecond, fold 和 tzinfo 属性性的 time 对象。 另请参见 time() 方法。

在 3.6 版更改: fold 值会被复制给返回的 time 对象。

`datetime.replace(year=self.year, month=self.month, day=self.day, hour=self.hour, minute=self.minute, second=self.second, microsecond=self.microsecond, tzinfo=self.tzinfo, * fold=0)`
返回一个具有同样属性值的 datetime，除非通过任何关键字参数为某些属性指定了新值。 请注意可以通过指定 tzinfo=None 来从一个感知型 datetime 创建一个简单型 datetime 而不必转换日期和时间数据。

3.6 新版功能: 增加了 fold 参数。

`datetime.astimezone(tz=None)`
返回一个具有新的 tzinfo 属性 tz 的 datetime 对象，并会调整日期和时间数据使得结果对应的 UTC 时间与 self 相同，但为 tz 时区的本地时间。

如果给出了 tz，则它必须是一个 tzinfo 子类的实例，并且其 utcoffset() 和 dst() 方法不可返回 None。 如果 self 为简单型，它会被假定为基于系统时区表示的时间。

如果调用时不传入参数 (或传入 tz=None) 则将假定目标时区为系统的本地时区。 转换后 datetime 实例的 .tzinfo 属性将被设为一个 timezone 实例，时区名称和时差值将从 OS 获取。

如果 self.tzinfo 为 tz，self.astimezone(tz) 等于 self: 不会对日期或时间数据进行调整。 否则结果为 tz 时区的本地时间，代表的 UTC 时间与 self 相同：在 astz = dt.astimezone(tz) 之后，astz - astz.utcoffset() 将具有与 dt - dt.utcoffset() 相同的日期和时间数据。

如果你只是想要附加一个时区对象 tz 到一个 datetime 对象 dt 而不调整日期和时间数据，请使用 dt.replace(tzinfo=tz)。 如果你只想从一个感知型 datetime 对象 dt 移除时区对象，请使用 dt.replace(tzinfo=None)。

请注意默认的 tzinfo.fromutc() 方法在 tzinfo 的子类中可以被重载，从而影响 astimezone() 的返回结果。 如果忽略出错的情况，astimezone() 的行为就类似于:

在 3.3 版更改: tz 现在可以被省略。

在 3.6 版更改: astimezone() 方法可以由简单型实例调用，这将假定其表示本地时间。

`datetime.utcoffset()`
如果 tzinfo 为 None，则返回 None，否则返回 self.tzinfo.utcoffset(self)，并且在后者不返回 None 或者一个幅度小于一天的 timedelta 对象时将引发异常。

在 3.7 版更改: UTC 时差不再限制为一个整数分钟值。

`datetime.dst()`
如果 tzinfo 为 None，则返回 None，否则返回 self.tzinfo.dst(self)，并且在后者不返回 None 或者一个幅度小于一天的 timedelta 对象时将引发异常。

在 3.7 版更改: DST 差值不再限制为一个整数分钟值。

`datetime.tzname()`
如果 tzinfo 为 None，则返回 None，否则返回 self.tzinfo.tzname(self)，如果后者不返回 None 或者一个字符串对象则将引发异常。

`datetime.timetuple()`
返回一个 time.struct_time，即 time.localtime() 所返回的类型。

d.timetuple() 等价于:

time.struct_time((d.year, d.month, d.day,
                  d.hour, d.minute, d.second,
                  d.weekday(), yday, dst))
其中 yday = d.toordinal() - date(d.year, 1, 1).toordinal() + 1 是日期在当前年份中的序号，起始序号 1 表示 1 月 1 日。 结果的 tm_isdst 旗标的设定会依据 dst() 方法：如果 tzinfo 为 None 或 dst() 返回 None，则 tm_isdst 将设为 -1；否则如果 dst() 返回一个非零值，则 tm_isdst 将设为 1；在其他情况下 tm_isdst 将设为 0。

`datetime.utctimetuple()`
如果 datetime 实例 d 为简单型，这类似于 d.timetuple()，不同之处在于 tm_isdst 会强制设为 0，无论 d.dst() 返回什么结果。 DST 对于 UTC 时间永远无效。

如果 d 为感知型， d 会通过减去 d.utcoffset() 来标准化为 UTC 时间，并返回该标准化时间所对应的 time.struct_time。 tm_isdst 会强制设为 0。 请注意如果 d.year 为 MINYEAR 或 MAXYEAR 并且 UTC 调整超出一年的边界则可能引发 OverflowError。

警告 由于简单型 datetime 对象会被许多 datetime 方法当作本地时间来处理，最好是使用感知型日期时间来表示 UTC 时间；因此，使用 utcfromtimetuple 可能会给出误导性的结果。 如果你有一个表示 UTC 的简单型 datetime，请使用 datetime.replace(tzinfo=timezone.utc) 将其改为感知型，这样你才能使用 datetime.timetuple()。

`datetime.toordinal()`
返回日期的预期格列高利历序号。 与 self.date().toordinal() 相同。

`datetime.timestamp()`
返回对应于 datetime 实例的 POSIX 时间戳。 此返回值是与 time.time() 返回值类似的 float 对象。

简单型 datetime 实例会假定为代表本地时间，并且此方法依赖于平台的 C mktime() 函数来执行转换。 由于在许多平台上 datetime 支持的范围比 mktime() 更广，对于极其遥远的过去或未来此方法可能引发 OverflowError。

对于感知型 datetime 实例，返回值的计算方式为:

(dt - datetime(1970, 1, 1, tzinfo=timezone.utc)).total_seconds()
3.3 新版功能.

在 3.6 版更改: timestamp() 方法使用 fold 属性来消除重复间隔中的时间歧义。

注解 没有一个方法能直接从表示 UTC 时间的简单型 datetime 实例获取 POSIX 时间戳。 如果你的应用程序使用此惯例并且你的系统时区不是设为 UTC，你可以通过提供 tzinfo=timezone.utc 来获取 POSIX 时间戳:
timestamp = dt.replace(tzinfo=timezone.utc).timestamp()
或者通过直接计算时间戳:

timestamp = (dt - datetime(1970, 1, 1)) / timedelta(seconds=1)
`datetime.weekday()`
返回一个整数代表星期几，星期一为 0，星期天为 6。 相当于 self.date().weekday()。 另请参阅 isoweekday()。

`datetime.isoweekday()`
返回一个整数代表星期几，星期一为 1，星期天为 7。 相当于 self.date().isoweekday()。 另请参阅 weekday(), isocalendar()。

`datetime.isocalendar()` 参考 `date.isocalendar()`

`datetime.isoformat(sep='T', timespec='auto')`
返回一个以 ISO 8601 格式表示的日期和时间字符串：

YYYY-MM-DDTHH:MM:SS.ffffff，如果 microsecond 不为 0

YYYY-MM-DDTHH:MM:SS，如果 microsecond 为 0

如果 utcoffset() 返回值不为 None，则添加一个字符串来给出 UTC 时差：

YYYY-MM-DDTHH:MM:SS.ffffff+HH:MM[:SS[.ffffff]]，如果 microsecond 不为 0

YYYY-MM-DDTHH:MM:SS+HH:MM[:SS[.ffffff]]，如果 microsecond 为 0

可选参数 sep (默认为 'T') 为单个分隔字符，会被放在结果的日期和时间两部分之间。 例如:

可选参数 timespec 要包含的额外时间组件值 (默认为 'auto')。它可以是以下值之一：

'auto': 如果 microsecond 为 0 则与 'seconds' 相同，否则与 'microseconds' 相同。

'hours': 以两个数码的 HH 格式 包含 hour。

'minutes': 以 HH:MM 格式包含 hour 和 minute。

'seconds': 以 HH:MM:SS 格式包含 hour, minute 和 second。

'milliseconds': 包含完整时间，但将秒值的小数部分截断至微秒。 格式为 HH:MM:SS.sss

'microseconds': 以 HH:MM:SS.ffffff 格式包含完整时间。

注解 排除掉的时间部分将被截断，而不是被舍入。
对于无效的 timespec 参数将引发 ValueError:

3.6 新版功能: 增加了 timespec 参数。

`datetime.__str__()`
对于 datetime 实例 d，str(d) 等价于 d.isoformat(' ')。

`datetime.ctime()`
返回一个表示日期和时间的字符串:

输出字符串将 并不 包括时区信息，无论输入的是感知型还是简单型。

d.ctime() 等效于:

time.ctime(time.mktime(d.timetuple()))
在原生 C ctime() 函数 (time.ctime() 会发起调用该函数，但 datetime.ctime() 则不会) 遵循 C 标准的平台上。

`datetime.strftime(format)`
返回一个由显式格式字符串所指明的代表日期和时间的字符串，要获取格式指令的完整列表，请参阅 strftime() 和 strptime() 的行为。

`datetime.__format__(format)`
与 datetime.strftime() 相同。 此方法使得为 datetime 对象指定以 格式化字符串字面值 表示的格式化字符串以及使用 str.format() 进行格式化成为可能。 要获取格式指令的完整列表，请参阅 strftime() 和 strptime() 的行为。

用法示例: datetime
使用 datetime 对象的例子：

## 1.4. datetime.time

&emsp;&emsp;一个独立于任何特定日期的理想化时间，它假设每一天都恰好等于 24*60*60 秒。 （这里没有“闰秒”的概念。） 包含属性: hour, minute, second, microsecond 和 tzinfo。

### 1.4.1. 类方法

&emsp;&emsp;一个 time 对象代表某日的（本地）时间，它独立于任何特定日期，并可通过 tzinfo 对象来调整。

- `class datetime.time(hour=0, minute=0, second=0, microsecond=0, tzinfo=None, *, fold=0)`
  - 所有参数都是可选的。 tzinfo 可以是 None，或者是一个 tzinfo 子类的实例。 其余的参数必须是在下面范围内的整数：
  - 0 <= hour < 24,0 <= minute < 60,0 <= second < 60,0 <= microsecond < 1000000,fold in [0, 1].
  - 如果给出一个此范围以外的参数，则会引发 ValueError。 所有参数值默认为 0，只有 tzinfo 默认为 None。
- `classmethod time.fromisoformat(time_string)`
  - 返回对应于 time.isoformat() 所提供的某种 time_string 格式的 time。 特别地，此函数支持以下格式的字符串：
  - HH[:MM[:SS[.fff[fff]]]][+HH:MM[:SS[.ffffff]]]
  - 警告 此方法 并不 支持解析任意 ISO 8601 字符串。 它的目的只是作为 time.isoformat() 的逆操作。

### 1.4.2. 类属性

- `time.min`
  - 早最的可表示 time, time(0, 0, 0, 0)。
- `time.max`
  - 最晚的可表示 time, time(23, 59, 59, 999999)。
- `time.resolution`
  - 两个不相等的 time 对象之间可能的最小间隔，timedelta(microseconds=1)，但是请注意 time 对象并不支持算术运算。

### 1.4.3. 实例属性

- `time.hour`
  - 取值范围是 range(24)。
- `time.minute`
  - 取值范围是 range(60)。
- `time.second`
  - 取值范围是 range(60)。
- `time.microsecond`
  - 取值范围是 range(1000000)。
- `time.tzinfo`
  - 作为 tzinfo 参数被传给 time 构造器的对象，如果没有传入值则为 None。
- `time.fold`
  - 取值范围是 [0, 1]。 用于在重复的时间段中消除边界时间歧义。 （当夏令时结束时回拨时钟或由于政治原因导致当明时区的 UTC 时差减少就会出现重复的时间段。） 取值 0 (1) 表示两个时刻早于（晚于）所代表的同一边界时间。

&emsp;&emsp;time 对象支持 time 与 time 的比较，当 a 时间在 b 之前时，则认为 a 小于 b。 如果比较的一方是简单型而另一方是感知型，则如果尝试进行顺序比较将引发 TypeError。 对于相等比较，简单型实例将永远不等于感知型实例。

&emsp;&emsp;如果两个比较方都是感知型，且具有相同的 tzinfo 属性，相同的 tzinfo 属性会被忽略并对基本时间值进行比较。 如果两个比较方都是感知型且具有不同的 tzinfo 属性，两个比较方将首先通过减去它们的 UTC 时差（从 self.utcoffset() 获取）来进行调整。 为了防止将混合类型比较回退为基于对象地址的默认比较，当 time 对象与不同类型的对象比较时，将会引发 TypeError，除非比较运算符是 == 或 !=。 在后两种情况下将分别返回 False 或 True。

### 1.4.4. 实例方法

- `time.replace(hour=self.hour, minute=self.minute, second=self.second, microsecond=self.microsecond, tzinfo=self.tzinfo, * fold=0)`
  - 返回一个具有同样属性值的 time，除非通过任何关键字参数指定了某些属性值。 请注意可以通过指定 tzinfo=None 从一个感知型 time 创建一个简单型 time，而不必转换时间数据。
- `time.isoformat(timespec='auto')`
  - 返回表示为下列 ISO 8601 格式之一的时间字符串：
    - HH:MM:SS.ffffff，如果 microsecond 不为 0
    - HH:MM:SS，如果 microsecond 为 0
    - HH:MM:SS.ffffff+HH:MM[:SS[.ffffff]]，如果 utcoffset() 不返回 None
    - HH:MM:SS+HH:MM[:SS[.ffffff]]，如果 microsecond 为 0 并且 utcoffset() 不返回 None
  - 可选参数 timespec 要包含的额外时间组件值 (默认为 'auto')。它可以是以下值之一：
    - 'auto': 如果 microsecond 为 0 则与 'seconds' 相同，否则与 'microseconds' 相同。
    - 'hours': 以两个数码的 HH 格式 包含 hour。
    - 'minutes': 以 HH:MM 格式包含 hour 和 minute。
    - 'seconds': 以 HH:MM:SS 格式包含 hour, minute 和 second。
    - 'milliseconds': 包含完整时间，但将秒值的小数部分截断至微秒。 格式为 HH:MM:SS.sss
    - 'microseconds': 以 HH:MM:SS.ffffff 格式包含完整时间。
  - 注解 排除掉的时间部分将被截断，而不是被舍入。
- `time.__str__()`
  - 对于时间对象 t, str(t) 等价于 t.isoformat()。
- `time.strftime(format)`
  - 返回一个由显式格式字符串所指明的代表时间的字符串。 要获取格式指令的完整列表，请参阅 strftime() 和 strptime() 的行为。
- `time.__format__(format)`
  - 与 time.strftime() 相同。 此方法使得为 time 对象指定以 格式化字符串字面值 表示的格式化字符串以及使用 str.format() 进行格式化成为可能。 要获取格式指令的完整列表，请参阅 strftime() 和 strptime() 的行为。
- `time.utcoffset()`
  - 如果 tzinfo 为 None，则返回 None，否则返回 self.tzinfo.utcoffset(None)，并且在后者不返回 None 或一个幅度小于一天的 a timedelta 对象时将引发- 异常。
  - 在 3.7 版更改: UTC 时差不再限制为一个整数分钟值。
- `time.dst()`
  - 如果 tzinfo 为 None，则返回 None，否则返回 self.tzinfo.dst(None)，并且在后者不返回 None 或者一个幅度小于一天的 timedelta 对象时将引发异常。
  - 在 3.7 版更改: DST 差值不再限制为一个整数分钟值。
- `time.tzname()`
  - 如果 tzinfo 为 None，则返回 None，否则返回 self.tzinfo.tzname(None)，如果后者不返回 None 或者一个字符串对象则将引发异常。

## 1.5. datetime.date

&emsp;&emsp;一个理想化的简单型日期，它假设当今的公历在过去和未来永远有效。 属性: year, month, and day。

### 1.5.1. 类方法

- `class datetime.date(year, month, day)`
  - 所有参数都是必要的。
  - 构造函数
- `classmethod date.today()`
  - 返回当前的本地日期。这等价于 date.fromtimestamp(time.time())。
- `classmethod date.fromtimestamp(timestamp)`
  - 返回对应于 POSIX 时间戳的当地时间
  - 这可能引发 OverflowError，如果时间戳数值超出所在平台 C localtime() 函数的支持范围的话，并且会在 localtime() 出错时引发 OSError。
  - 通常该数值会被限制在 1970 年至 2038 年之间。 请注意在时间戳概念包含闰秒的非 POSIX 系统上，闰秒会被 fromtimestamp() 所忽略。
- `classmethod date.fromordinal(ordinal)`
  - 返回对应于预期格列高利历序号的日期，其中公元 1 年 1 月 1 晶的序号为 1。
- `classmethod date.fromisoformat(date_string)`
  - 返回一个对应于以 YYYY-MM-DD 格式给出的 date_string 的 date 对象
- `classmethod date.fromisocalendar(year, week, day)`
  - 返回指定 year, week 和 day 所对应 ISO 历法日期的 date。 这是函数 date.isocalendar() 的逆操作。
  - 3.8 新版功能.

### 1.5.2. 实例方法

- `date.replace(year=self.year, month=self.month, day=self.day)`
  - 返回一个具有同样值的日期，除非通过任何关键字参数给出了某些形参的新值。
- `date.timetuple()`
  - 返回一个 time.struct_time，即 time.localtime() 所返回的类型。
  - hours, minutes 和 seconds 值均为 0，且 DST 旗标值为 -1。
  - d.timetuple() 等价于:time.struct_time((d.year, d.month, d.day, 0, 0, 0, d.weekday(), yday, -1))
  - 其中 yday = d.toordinal() - date(d.year, 1, 1).toordinal() + 1 是当前年份中的日期序号，1 月 1 日的序号为 1。
- `date.toordinal()`
  - 返回日期的预期格列高利历序号，其中公元 1 年 1 月 1 日的序号为 1。 对于任意 date 对象 d，date.fromordinal(d.toordinal()) == d。
- `date.weekday()`
  - 返回一个整数代表星期几，星期一为0，星期天为6。例如， date(2002, 12, 4).weekday() == 2，表示的是星期三。参阅 isoweekday()。
- `date.isoweekday()`
  - 返回一个整数代表星期几，星期一为1，星期天为7。例如：date(2002, 12, 4).isoweekday() == 3,表示星期三。参见 weekday(), isocalendar()。
- `date.isocalendar()`
  - 返回一个由三部分组成的 named tuple 对象: year, week 和 weekday。
  - ISO 历法是一种被广泛使用的格列高利历。
  - ISO 年由 52 或 53 个完整星期构成，每个星期开始于星期一结束于星期日。 一个 ISO 年的第一个星期就是（格列高利）历法的一年中第一个包含星期四的星期。 - 这被称为 1 号星期，这个星期四所在的 ISO 年与其所在的格列高利年相同。
  - 例如，2004 年的第一天是星期四，因此 ISO 2004 年的第一个星期开始于 2003 年 12 月 29 日星期一，结束于 2004 年 1 月 4 日星期日:
- `date.isoformat()`
  - 返回一个以 ISO 8601 格式 YYYY-MM-DD 来表示日期的字符串:
- `date.__str__()`
  - 对于日期对象 d, str(d) 等价于 d.isoformat() 。
- `date.ctime()`
  - 返回一个表示日期的字符串
  - d.ctime() 等效于:time.ctime(time.mktime(d.timetuple()))
  - 在原生 C ctime() 函数 (time.ctime() 会发起调用该函数，但 date.ctime() 则不会) 遵循 C 标准的平台上。
- `date.strftime(format)`
  - 返回一个由显式格式字符串所指明的代表日期的字符串。 表示时、分或秒的格式代码值将为 0。 要获取格式指令的完整列表请参阅 strftime() 和 strptime() 的- 行为。
- `date.__format__(format`)
  - 与 date.strftime() 相同。 此方法使得为 date 对象指定以 格式化字符串字面值 表示的格式化字符串以及使用 str.format() 进行格式化成为可能。 要获取格式指令的完整列表，请参阅 strftime() 和 strptime() 的行为。

### 1.5.3. 类属性

- `date.min`
  - 最小的日期 date(MINYEAR, 1, 1) 。
- `date.max`
  - 最大的日期 ，date(MAXYEAR, 12, 31)。
- `date.resolution`
  - 两个日期对象的最小间隔，timedelta(days=1)。

### 1.5.4. 实例属性

- `date.year`
  - 在 MINYEAR 和 MAXYEAR 之间，包含边界。
- `date.month`
  - 1 至 12（含）
- `date.day`
  - 返回1到指定年月的天数间的数字。

### 1.5.5. 支持的运算

- `date2 = date1 + timedelta`
  - date2 等于从 date1 减去 timedelta.days 天.
  - 如果 timedelta.days > 0 则 date2 将在时间线上前进，如果 timedelta.days < 0 则将后退。 操作完成后 date2 - date1 == timedelta.days。 timedelta.seconds 和 timedelta.microseconds 会被忽略。 如果 date2.year 将小于 MINYEAR 或大于 MAXYEAR 则会引发 OverflowError。
- `date2 = date1 - timedelta`
  - 计算 date2 的值使得 date2 + timedelta == date1。
  - timedelta.seconds 和 timedelta.microseconds 会被忽略。
- `timedelta = date1 - date2`
  - 此值完全精确且不会溢出。 操作完成后 timedelta.seconds 和 timedelta.microseconds 均为 0，并且 date2 + timedelta == date1。
- `date1 < date2`
  - 如果 date1 的时间在 date2 之前则认为 date1 小于 date2 。 (4)

## 1.6. datetime.tzinfo

&emsp;&emsp;一个描述时区信息对象的抽象基类。 用来给 datetime 和 time 类提供自定义的时间调整概念（例如处理时区和/或夏令时）。

### 1.6.1. datetime.timezone

&emsp;&emsp;一个实现了 tzinfo 抽象基类的子类，用于表示相对于 世界标准时间（UTC）的偏移量。

## 1.7. strftime() 和 strptime() 的行为

## 1.8. 参考资料

1. [官网](https://docs.python.org/zh-cn/3/library/datetime.html)
