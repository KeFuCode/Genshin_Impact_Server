Genshin_Impact_Server 原神服务端/后端
# 0 文件结构
1. bin 文件夹
   - csvs 文件夹：放入配置表
   - 编译后的执行文件
2. csv 文件夹  
   csv 格式的配置数据
3. game 文件夹
   - manage_banword.go
   - mod_card.go
   - mod_icon.go
   - mod_player.go
   - mod_role.go
   - mod_uniquetask.go
   - player.go
4. utils 文件夹
   - csvutils.go
5. go.mod
   包依赖管理
6. main.go
   程序的入口

# 1 基础模块
## 1.1 基础模块
玩家从客户端登陆后，首先拿到的就是基础信息。
这个模块对应数据库中的一张表

UID 对每个玩家是唯一确定的

基础模块内容：
1. UID
2. 头像，名片
3. 签名
4. 名字
5. 冒险等级 冒险阅读
6. 世界等级 冷却时间：时间用 int64 存储
7. 生日
8. 展示阵容 展示名片

> 隐藏内容
1. 账号状态：使用 int 保存（不用 bool ），方便后续扩展。 
2. 管理员状态

## 1.2 模块关系
Player 作为最上层模块，接收到 Client 发送的消息后，调用 ModPlayer 内置方法。

## 1.3 名字验证
描述：名字验证
1. 英文名：字符格式多样，一般不做处理。
2. 中文名：违规中文名组合多种多样，防不胜防。直接写验证条件，无法穷尽，对服务器要求高。
   - 做法1：调用 HTTP 地址接口，进行外部验证。
     - 优点：对服务器压力小。
     - 缺点： 依赖外部验证服务器的安全性、稳定性。
   - 做法2（常用）：设置违禁词库，并提供更新方法。

例子：ManageBanWord 名字验证

描述：管理类与玩家类区别
玩家连接客户端之后，玩家线程被动创建。公共管理类，需要 server 启动后主动创建。
例子：单例模式实现管理类 管理类与玩家类区别

## 1.4 定时器功能
增加一个定时器，实现扩充禁用词。

描述：定时器作用
定时器给行为赋予了时间维度，每隔相同时间，就会触发一次行为。
例子：Run() 定时器作用
```go
func (self *ManageBanWord) Run() {
	triker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-triker.C:
			fmt.Println(time.Now().Unix())
		}
	}
}
```

## 1.5 启动协议 
将定时器 Run() 以协程 `go Run()` 的形式启动，可以不干扰主程序的运行。
否则程序会从 Run() 进入死循环，无法进入主程序后续逻辑。

## 1.6 配置基础词库  
一般策划将 EXCEL 表提交给程序，程序将 EXCEL 表转换为 CSV/Json 格式的内容。
Json 格式更像一个 map ，CSV 格式更像数组。两者可以互相转化，关键取决于服务器内部使用什么数据结构。

配置文件会在程序运行之后，初始化至内存中。init() 函数在初次调用 package 是，会自动执行。

## 1.7 人物等级和经验
### 1.7.1 配置原始数据
utils 中放入处理原始数据的脚本代码，方便将 CSV 文件载入内存。
### 1.7.2 人物升级方法
internal function: `func (self *ModPlayer) AddExp(exp int)` 
不能从客户端调用，只有服务器可以调用，当玩家进行某些行为（杀怪、副本、探索等），服务器根据情况，为玩家增加经验值。

## 1.8 完成任务
完成突破任务，人物等级才能继续提升。
map 查找效率很高，但是多线程中如果不加锁，很容易崩溃。

添加 ModUniqueTask , 用于玩家突破升级的验证条件。

> 完成任务的功能，是服务器内置功能，为了防止玩家在客户端反复刷资源。
## 1.9 多线程锁和 map
为了保证数据的安全性，在内存中的数据，同一时间不能多个线程共同操作。
当发生多个线程同时读写内存中同样数据时，主线程会发生异常，产生的结果是服务器崩溃，属于严重的安全事故。
给读/写操作加锁，虽然保证了数据的安全性，但是对服务器性能影响很大（3倍左右的差距）。

> 锁的基础理论与实际业务中的冲突  

map 的时间复杂度 O(1) ， 由于 map 的性能好，实际生产中经常使用 map ， 那么如何保证数据安全的同时，不降低服务器性能？
1. 改变设计模式，将多个线程降为一个线程，将多把锁降为一个锁。
   游戏中发放奖励，通常是通过邮件完成，而不是直接向玩家的背包中添加。直接向玩家背包中添加，若玩家在线，则背包中的数据面临同时读写，需要大量的锁才能保证背包数据安全。改变为邮件之后，是玩家通过领取邮件，自行将奖励写入背包，因此只需要为邮件系统加锁，不必为每个背包加锁。
2. 对于只读的数据，即使使用map，也不会影响数据安全性。
   配置表中的 `var ConfigUniqueTaskMap map[int]*ConfigUniqueTask` 是 map 类型，但是配置表不需要修改，只有读操作。不管多少线程去读，也不会导致数据不一致。 
## 1.10 世界等级
实现主动降低世界等级，恢复原来世界等级接口。两者都由玩家触发，是外部接口。
## 1.11 设置生日
Golang 中 `time.Now()` 获取的是系统本地时区对应的时间，可以直接使用，不需要额外时区转换。
## 1.12 名片展示
基础模块简化为三个部分：展示阵容，展示名片和封禁状态。
基础模块做完之后，就可以做登录模块。

客户端容易被修改，服务端需要验证客户端发送过来的多参数是否合法。只有合法的数据才被接收。

对于 `map[int]int` 可以通过 `v, ok := map[int]` 访问 map 中的值，其中 v 是获取的值，ok 是 bool 值，表示 map 中是否存在该键对应的值（存在返回 true，否则为 false ）。
## 1.13 团队展示
bool 值在实际使用中较少，true/false 在数据库中存储时，有时是 true/false ，有时是 0/1 ，为了统一和扩展，一般都使用 int 0/1 来表示 bool 值。

实现思路与名片类似。

限定名片/团队展示尺寸最大为 9 ，若大于 9 ，说明是修改客户端的违法数据，直接返回。
## 1.14 封禁和 GM
`Prohibit` 使用 int 类型存储，可以保存多种状态，如：封禁一周、封禁一个月、封禁一年等。
游戏中最多封禁到 203x 年，参考 int64 和 int 取值范围，使用的应该是 int 类型。

封禁和 GM 功能都是由服务器触发，不被外部调用，因此是 `internal function` 。
# 2 背包模块
资源管理，负责各模块间资源调配。
将 icon card 等道具与 id 一一对应，便于资源分配与管理。

模块功能：
1. 物品识别
2. 物品增加
3. 物品消耗
4. 物品使用
5. 角色模块 --> 头像模块

背包模块为什么不称为资源管理模块？  
背包模块只处理部分物品，自己不方便处理、或不愿意处理的物品，交给其他模块处理。

## 2.1 物品识别
根据 itemId 识别出物品名称。
增加 csv.item.go 保存 item 配置信息，mod_bag 作为背包模块。

## 2.2 补全头像模块
增加 icon 配置信息，补全头像模块。
根据物品类别，将头像信息加入背包，玩家可以设置已获得头像。

## 2.3 名片模块和角色好感度
增加 card 配置信息，补全名片模块，与头像模块类似。

## 2.4 背包模块杂物功能
其他模块都不负责的物品，由背包模块进行处理。
### 2.4.1 移除物品
- 玩家移除物品
   玩家移除物品，为了保证数据安全，需要两次判断。第一次是调用函数，直接判断；第二次是执行移除物品操作前判断。不允许赊账。
- 管理员移除物品
  为满足恶意退款情况，可以将物品数量设为负数。

## 2.5 角色模块
玩家获得角色，根据角色数量，自动添加进背包。

## 2.6 简单测试功能
游戏行业中的交互，大多是消息驱动。客户端将想做的动作发送给服务器，服务器验证数据后，修改数据，就完成了相应的行为。

借助命令行，实现简单测试功能。添加测试模块，用于后续测试。

黑盒测试：从用户角度，判断游戏逻辑是否正常运行。

## 2.7 头像模块
头像为何要单独建一个表，不直接放入角色表中？
若头像放入角色表，后续与角色有关的内容，都要放入角色表。容易造成角色表内容过于复杂。因此头像需要单独成表。
有专门的模块，最好就需要有专门的表与之对应。

用 IconId 作为 key 值，获取 role 之后，需要遍历所有头像才能找到与 RoleId 匹配的 Check 值，时间复杂度过高。
解决方案：用空间换时间，增加一个以 RoleId 为键值的 map。

## 2.8 名片模块
获得 role 时，自动获得 role 对应的 card 。默认设置友好度为 10 ，后续实现该功能。
实现思路与 `2.7 头像模块` 相同。

## 2.9 武器模块
武器也需要单独模块进行管理。每个武器有自己的专属 id ， 各武器之间不可以相互堆叠。
武器支持批量获取，有总量最大上限。

## 2.10 圣遗物模块
实现过程参照武器模块，只有最大数量上限的区别。

## 2.11 烹饪技能背包
烹饪技能书和烹饪的食物，都归背包模块的杂物部分统一管理。学习、学会的烹饪技能，由烹饪模块单独管理。

获得的食谱，添加进背包中，可以使用。使用食谱后，食谱消失，获得食谱对应的技能。

## 2.12 家园模块
家园模块功能较复杂，一方面获得的家具会添加进背包，另一方面家具的摆放（位置、数量、地点等）需要家园模块单独管理。
在设计家园模块时，为了响应的速度，使用空间换时间的思路。

## 2.13 总结
已实现：人物信息查询和设置（角色、头像、名片），增加物品，移除物品，使用物品。
背包模块中的子模块实现相似，逻辑较为简单。家园模块涉及多个 map 的相互调用，实现起来有一定难度。

# 3 掉落模块
需求分析：
1. 保底设计
2. 大数据测试（为策划提供决策数据）
3. 配套测试工具
4. up 角色池子
5. 仓检

## 3.1 表格配置
当遇到抽卡 up 时，合理配置 Drop.csv 表格，使数据维护更加方便。

难点： `csv_check.go`  中根据 drop 的配置信息，生成新的数据结构。
```go
func MakeDropGroupMap() {
	configDropGroupMap := make(map[int]*DropGroup)
	for _, v := range ConfigDropSlice {
		dropGroup, ok := configDropGroupMap[v.DropId]
		if !ok {
			dropGroup = new(DropGroup)
			dropGroup.DropId = v.DropId
			configDropGroupMap[v.DropId] = dropGroup
		}
		dropGroup.WeightAll += v.Weight
		dropGroup.DropConfigs = append(dropGroup.DropConfigs, v)
	}
}
```