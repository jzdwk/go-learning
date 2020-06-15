package structural

/*
组合模式（Composite Pattern），又叫部分整体模式，是用于把一组相似的对象当作一个单一的对象。
组合模式依据树形结构来组合对象，用来表示部分以及整体层次。
这种类型的设计模式属于结构型模式，它创建了对象组的树形结构。

应用场景：
	组合模式描述的是一种递归结构的struct，比如单位的组织结构、员工层级关系等。以组织结构为例，
	对于每一级组织，都有类似增加子机构、修改本机名称等功能
关键代码：
	递归结构的struct
*/

type organization interface {
	name() string
	addSub(org organization) organization
	delSub(name string) organization
	print() string
}

//递归的struct层次
type myOrg struct {
	msg    string
	subOrg []organization
}

func (o *myOrg) name() string {
	return o.msg
}
func (o *myOrg) addSub(sub organization) organization {
	o.subOrg = append(o.subOrg, sub)
	return o
}
func (o *myOrg) delSub(name string) organization {
	for i, v := range o.subOrg {
		if v.name() == name {
			o.subOrg = o.subOrg[:i+copy(o.subOrg[i:], o.subOrg[i+1:])]
		}
	}
	return o
}
func (o *myOrg) print() string {
	org := "name: " + o.name()
	//只打印向下一层
	for _, o := range o.subOrg {
		org += "name" + o.name()
	}
	return org
}

func NewMyOrg(orgName string) organization {
	return &myOrg{orgName, []organization{}}
}
