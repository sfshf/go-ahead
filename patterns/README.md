参考资料：
1. [《设计模式之禅（第2版）》秦小波](https://www.amazon.cn/gp/product/B00INI842W?psc=1)

2. [github.com/senghoo/golang-design-pattern](https://github.com/senghoo/golang-design-pattern)

3. [github.com/me115/design_patterns](https://github.com/me115/design_patterns)

**【一】六大设计原则**

***（1）单一职责原则（Single Resposibility Principle）***

SRP原话解释：There should never be more than one reason for a class to change.

单一职责原则要求一个接口或类只有一个原因引起变化，也就是一个接口或类只有一个职责，它就负责一件事情。

***（2）里氏替换原则（Liskv Substitution Principle）***

LSP有两种定义：
 · 第一种定义，也是最正宗的定义：
   If for each object o1 of type S there is an object o2 of type T such that for all programs P defined in terms of T, the behavior of P is unchanged when o1 is substituted for o2 then S is a subtype of T.    
 · 第二种定义：
   Functions that user pointers or references to base classes must be able to use objects of derived classes without knowing it.

第二个定义是最清晰明确的，通俗点讲，只要父类能出现的地方子类就可以出现，而且替换为子类也不会产生任何错误或异常，使用者可能根本就不需要知道父类还是子类。但是反过来就不行了，有子类出现的地方，父类未必就能适应。

***（3）依赖倒置原则（Dependence Inversion Principle）***

DIP原始定义：
    High level modules should not depend upon low level modules. Both should depend upon abstractions. Abstractions should not depend upon details. Details should depend upon abstractions.
翻译过来，包含三层含义：
    · 高层模块不应该依赖低层模块，两者都应该依赖其抽象；
    · 抽象不应该依赖细节；
    · 细节应该依赖抽象。

***（4）接口隔离原则（Interface Segregation Principle）***

ISP有两种定义：
    · Clients should not be forced to depend upon interfaces that they don't use.（客户端不应该依赖它不需要的接口。）
    · The dependency of one class to another one should depend upon the smallest possible interface.（类间的依赖关系应该建立在最小的接口上。）

接口隔离原则是对接口进行规范约束，其包含以下4层含义：
    · 接口要尽量小；
    · 接口要高内聚；
    · 定制服务；
    · 接口设计是有限度的。


***（5）迪米特原则（Law Of Demeter）|最少知识原则（Least Knowledge Principle）***

迪米特法则描述的是一个规则：一个对象应该对其他对象有最少的了解。

迪米特法则的核心观念就是类间解耦，弱耦合，只有弱耦合了以后，类的复用率才可以提高。其要求的结果就是产生了大量的中转或跳转类，导致系统的复杂性提高，同时也为维护带来了难度。程序员在采用迪米特法则时需要反复权衡，既做到让结构清晰，又做到高内聚低耦合。

***（6）开闭原则（Open Closed Principle）***

开闭原则的定义：
    Software entities like classes, modules and functions should be open for extension but closed for modifications.（一个软件实体如类、模块和函数应该对扩展开放，对修改关闭。）

前五个原则就是指导设计的工具和方法，而开闭原则才是精神领袖。
开闭原则是一个口号，那么怎么把这个口号应用到实际工作中呢？
    · 抽象约束；
    · 元数据（metadata）控制模块行为；典型的元数据控制模块行为的例子有控制反转(IOC)；
    · 制定项目章程；
    · 封装变化。


*****

软件设计最大的难题就是应对需求的变化，但是纷繁复杂的需求变化又是不可预料的。我们要为不可预料的事情做好准备，这本身就是一件非常痛苦的事情，但是大师们还是给我们提出了非常好的经典6大设计原则以及经典23个设计模式（除解释器模式外）来“封装”未来的变化。

六大设计原则的英文首字母联合起来就是SOLID，其含义就是把这6个原则结合使用的好处：建立稳定、灵活、健壮的设计，而开闭原则又是重中之重，是最基础的原则，是其他5大原则的精神领袖。

*****
