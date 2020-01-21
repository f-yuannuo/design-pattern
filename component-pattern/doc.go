package component_pattern

/*
==================组合模式===================
	1.定义：允许你将对象组成树形结构来表现"整体/部分"层次结构。组合能让客户以一致的方式处理个别对象以及对象组合
	2.本例子的话以菜单为例，菜单有子菜单。在menuComponent.go文件里面定义了组合模式的所有接口，然后我们用一个空结构体去实现了他
	menu.go继承了这个空结构体并实现了MenuComponenter的接口。如我我们要实现树形结构就需要像这样在自己结构体的属性中包含自己，但是我们现在用的是接口
	menuItem.go就是具体的菜单，有价格有菜名这样一个菜单的树形结构就完成了。
*/