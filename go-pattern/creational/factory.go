package creational

/*
wiki: 工厂方法模式（英语：Factory method pattern）是一种实现了“工厂”概念的面向对象设计模式。
就像其他创建型模式一样，它也是处理在不指定对象具体类型的情况下创建对象的问题。
工厂方法模式的实质是“定义一个创建对象的接口，但让实现这个接口的类来决定实例化哪个类。
工厂方法让类的实例化推迟到子类中进行。”

例子： https://www.zhihu.com/question/20367734

*/

/*
场景：产品接口为mouse，定义了鼠标类产品。具体的鼠标生产包括了联想鼠标和HP鼠标
*/

//鼠标接口
type mouse interface {
	doCreate(msg string) (out string, err error)
}

//实现接口的结构1 联想鼠标
type lenovoMouse struct {
	msg string
}

func (imp1 *lenovoMouse) doCreate(msg string) (string, error) {
	return imp1.msg + msg, nil
}

//实现接口的结构2 惠普鼠标
type hpMouse struct {
	msg string
}

func (imp2 *hpMouse) doCreate(msg string) (string, error) {
	return imp2.msg + msg, nil
}

/*
1.简单工厂模式，并不属于工厂方法

- 一个产品接口mouse以及不同的实现lenovn/hp，一个工厂方法NewSimpleFactory
- NewSimpleFactory通过入参和switch-case逻辑，返回不同的mouse实例
- 根据不同实例调用mouse接口的方法完成后续逻辑

缺点： 没有将变化的部分拆分(switch-case)
*/

func NewSimpleFactory(storeType string) mouse {
	switch storeType {
	case "lenovo":
		return &lenovoMouse{"lenovo:"}
	case "hp":
		return &hpMouse{"hp:"}
	default:
		return nil
	}
}

/*
2. 工厂方法，依赖工厂接口。
我们可以通过实现工厂接口，创建多种工厂。
将对象创建由一个对象负责所有具体类的实例化，变成由一群子类来负责对具体类的实例化。
将过程解耦。工厂方法需要:

- 工厂接口 factory
- 工厂结构体 lenovo/hp factory
- 产品接口 mouse
- 产品结构体 lenovo/hp factory

优点：
	通过factory接口的具体实现，将变化的产品创建从参数+switch-case中分离，通过创建不同的factory去执行不同产品的创建。
比如新增华硕产品，只需要新增华硕factory&mouse

缺点：
	每个factory的具体实现，都可以创建多种产品，从而对应多个产品的实例，这些products在不同的factory中可能是不同的，从而导致'类爆炸'。
比如增加了华硕，就要增加对应的factory&mouse实现
*/

type factory interface {
	createMouse(types string) mouse
}

type lenovoFactory struct {
}

func (f *lenovoFactory) createMouse(types string) mouse {
	switch types {
	case "lenovoMouse1":
		//do sth
		return &lenovoMouse{"lenovoFactory produces lenovoMouse1"}
	//any case...
	default:
		return nil
	}
}

type hpFactory struct {
}

func (f *hpFactory) createMouse(types string) mouse {
	switch types {
	case "hpMouse1":
		//do sth
		return &hpMouse{"hpFactory  produces hpMouse1"}
	//any case...
	default:
		return nil
	}
}

func NewLenovoFactory() factory {
	return &lenovoFactory{}
}

func NewHpFactory() factory {
	return &hpFactory{}
}

/*
3. 抽象工厂模式。

工厂模式和简单工厂模式都是针对单一一种产品设计的，比如上边的mouse接口，只是具体实现上有lenovo和hp.
之前的方式能够完成联想/jp鼠标的生产场景（定义为纵向扩展，两种鼠标实例都继承了鼠标接口，增加一个厂家，只需要实现mouse）。
然而，当产品的种类不只一种，比如除了鼠标，又有键盘，则工厂模式就要扩展为抽象工厂。

优点：
	和工厂模式一样，解决了同类对象，不同实现的扩展，只是抽象工厂支持的同类对象不止一种
缺点：
	所有的工厂模式，当产品类别扩充（比如除了鼠标，键盘，新增了网线，则所有接口和实现都要修改）

*/

//增加一种产品keyboard
type keyboard interface {
	doCreate(msg string) (out string, err error)
}
type lenovoKeyboard struct {
	msg string
}

func (k *lenovoKeyboard) doCreate(msg string) (out string, err error) {
	return k.msg + msg, nil
}

type hpKeyboard struct {
	msg string
}

func (k *hpKeyboard) doCreate(msg string) (out string, err error) {
	return k.msg + msg, nil
}

//抽象工厂的定义，根据产品类别区分，是工厂的工厂
type pcAbFactory interface {
	//pd工厂生产mouse
	createMouse() mouse
	//keyboard
	createKeyboard() keyboard
}

//hp的实现
type hpAbFactory struct {
}

func (f *hpAbFactory) createMouse() mouse {
	return &hpMouse{msg: "this is hp mouse"}
}

func (f *hpAbFactory) createKeyboard() keyboard {
	return &hpKeyboard{msg: "this is hp keyboard"}
}

//lenovo的实现
type lenovoAbFactory struct {
}

func (f *lenovoAbFactory) createMouse() mouse {
	return &lenovoMouse{msg: "this is lenovo mouse"}
}

func (f *lenovoAbFactory) createKeyboard() keyboard {
	return &lenovoKeyboard{msg: "this is lenovo keyboard"}
}

/*
总结:
简单工厂 通过switch-case将对象创建逻辑封装。
工厂模式 围绕一个超级工厂创建其他工厂，将产品具体实现的变化（扩展）抽离，每个具体产品hp/lenovo mouse对应一个工厂hp/lenovo factory。当产品扩展，增加工厂和产品的实现(华硕)。
抽象工厂 和工厂模式的区别在于，要生产的产品有多种(除了mouse，还有keyboard)。因此，需要针对不同种类产品定义接口的方法

所有的工厂模式，能够进行纵向的扩展，比如定义好了产品种类为mouse/keyboard，这时增加一个华硕品牌。
但是对产品的横向扩展封闭。当产品类别扩充，比如除了鼠标，键盘，新增了网线，则所有接口和实现都要修改。

因此适用的场景为：
	产品的种类固定（api接口只有crud+list），产品的实现可扩展（多个api）

*/
