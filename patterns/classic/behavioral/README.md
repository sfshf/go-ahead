
[***职责链模式（Chain of Responsibility Pattern）***](ChainOfResponsibility)

定义:
    Avoid coupling the sender of a request to its receiver by giving more than one object a chance to handle the request. Chain the receiving objects and pass the request along the chain until an object handles it.（使多个对象都有机会处理请求，从而避免了请求的发送者和接受者之间的耦合关系。将这些对象连成一条链，并沿着这条链传递该请求，直到有对象处理它为止。）

[***命令模式（Command Pattern）***](Command)

定义：
    Encapsulate a request as an object, thereby letting you parameterize clients with different requests, queue or log requests and support undoable operations.（将一个请求封装成一个对象，从而让你使用不同的请求把客户端参数化，对请求排队或者记录请求日志，可以提供命令的撤销和恢复功能。）

[***解释器模式（Interpreter Pattern）***](Interpreter)

解释器模式是一种按照规定语法进行解析的方案，在现在项目中使用较少，如需解释器工具可以网上找优秀的工具。

定义：
    Given a language, define a representation for its grammar along with an interpreter that uses the representation to interpret sentences in the language.（给定一门语言，定义它文法的一种表示，并定义一个解释器，该解释器使用该表示来解释语言中句子。）

[***迭代器模式（Iterator Pattern）***](Iterator)

定义：
    Provide a way to access the elements of an aggregate object sequentially without exposing its underlying representation.（它提供一种方法访问一个容器对象中各个元素，而又不需暴露该对象的内部细节。）

[***中介者模式（Mediator Pattern）***](Mediator)

定义：
    Define an object that encapsulates how a set of objects interact. Mediator promotes loose coupling by keeping objects from referring to each other explicitly, and it lets you vary their interaction independently.（用一个中介对象封装一系列的对象交互，中介者使各对象不需要显示地相互作用，从而使其耦合松散，而且可以独立地改变它们之间的交互。）

[***备忘录模式（Memento Pattern）***](Memento)

备忘录模式提供了一种弥补真实世界缺陷的方法，让“后悔药”在程序的世界中真实可行。

定义：
    Without violating encapsulation, capture and externalize an object's internal state so that the object can be restored to this state later.（在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样之后就可将该对象恢复到原先保存的状态。）

[***观察者模式（Observer Pattern）***](Observer)

观察者模式也叫做发布订阅模式（Publish/Subscribe）。

定义：
    Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.（定义对象间一种一对多的依赖关系，使得每当一个对象改变状态，则所有依赖于它的对象都会得到通知并自动更新。）

[***状态模式（State Pattern）***](State)

定义：
    Allow an object to alter its behavior when its internal state changes. The object will appear to change its class.（当一个对象内在状态改变时允许其改变行为，这个对象看起来像改变了其类。）

状态模式的核心是封装，状态的变更引起了行为的变更，从外部看起来就好像这个对象对应的类发生了改变一样。

[***策略模式（Strategy Pattern）***](Strategy)

策略模式是一种比较简单的模式，也叫做政策模式（Policy Pattern）。

定义：
    Define a family of algorithms, encapsulate each one, and make them interchangeable.（定义一组算法，将每个算法都封装起来，并且使它们之间可以互换。）

[***访问者模式（Visitor Pattern）***](Visitor)

定义：
    Represent an operation to be performed on the elements of an object structure. Visitor lets you define a new operation without changing the classes of the elements on which it operats.（封装一些作用于某种数据结构中的个元素的操作，它可以在不改变数据结构的前提下定义作用于这些元素的新的操作。）

[***模板方法模式（Template Method Pattern）***](TemplateMethod)

定义：
    Define the skeleton of an algorithm in an operation, deferring some steps to subclasses. Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.（定义一个操作中的算法的框架，而将一些步骤延迟到子类中。使得子类可以不改变一个算法的结构可重定义该算法的某些特定步骤。）
