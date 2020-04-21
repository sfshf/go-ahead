
[***代理模式（Proxy Pattern）***](Proxy)

定义：
    Provide a surrogate or placeholder for another object to control access to it.（为其他对象提供一种代理控制对这个对象的访问。）

[***适配器模式（Adapter Pattern）***](Adapter)

定义：
    Convert the interface of a class into another interface clients expect. Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.（将一个类的接口变换成客户端所期待的另一种接口，从而使原本因接口不匹配而无法在一起工作的两个类能够在一起工作。）

适配器模式又叫做变压器模式，也叫作包装模式（Wrapper），但是包装模式可不止一个，还包括装饰模式。

[***桥接模式（Bridge Pattern）***](Bridge)

桥梁模式也叫作桥接模式。

定义：
    Decouple an abstraction from its implementation so that the two can vary independently.（将抽象和实现解耦，使得两者可以独立地变化。）

桥梁模式的重点是在“解耦”上，如何让它们两者解耦是我们要了解的重点。

[***组合模式（Composite Pattern）***](Composite)

组合模式也叫合成模式，有时也叫作“部分-整体模式”（Part-Whole），主要是用来描述部分与整体的关系。

定义：
    Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.（将对象组合成树形结构以表示“部分-整体”的层次结构，使得用户对单个对象和组合对象的使用具有一致性。）


[***装饰模式（Decorator Pattern）***](Decorator)

定义：
    Attach additional responsibilities to an object dynamically keeping the same interface. Decorators provide a flexible alternative to subclassing for extending functionality.（动态地给一个对象添加一些额外的职责。就增加功能来说，装饰模式相比生成子类更为灵活。）

[***外观模式（Facade Pattern）***](Facade)

门面模式也叫作外观模式。

定义：
    Provide a unified interface to a set of interfaces in a subsystem. Facade defines a higher-level interface that makes the subsystem easier to use.（要求一个子系统的外部与其内部的通信必须通过一个统一的对象进行。门面模式提供一个高层次的接口，使得子系统更易于使用。）

门面模式注重“统一的对象”，也就是提供一个访问子系统的接口，除了这个接口不允许有任何访问子系统的行为发生。

[***享元模式（Flyweight Pattern）***](Flyweight)

享元模式是池技术的重要实现方式。

定义：
    Use sharing to support large numbers of fine-grained objects efficiently.（使用共享对象可有效地支持大量的细粒度的对象。）

享元模式的定义为我们提出了两个要求：细粒度的对象和共享对象。我们知道分配太多的对象到应用程序中将有损程序的性能，同时还容易造成内存溢出，那怎么避免呢？就是享元模式提到的共享技术。
