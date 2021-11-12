package types

//type CorePool struct {
//	P            sync.Pool
//	registerList []Register
//}
//
//func (c *CorePool) RegisterList(l []Register) {
//	c.registerList = l
//}
//
//func (c *CorePool) Set(ctx *Core) {
//	log.Println("ctx pool: ", c)
//	for _, v := range c.registerList {
//		v.Set(ctx)
//	}
//}
//
//func (c *CorePool) Unset(ctx *Core) {
//	for _, v := range c.registerList {
//		v.Unset(ctx)
//	}
//}
//
//func (c *CorePool) Reload() {
//	ctx := c.GetObj()
//	for _, v := range c.registerList {
//		v.Reload(ctx)
//	}
//}
//
//// 从pool获取对象后进行初始化
//func (c *CorePool) GetInit() {
//	// Todo 获取pool后的初始化
//}
//
//func (c *CorePool) ReUse(ctx *Core) {
//	// tracer收尾 防止有的trace 没有end
//	ctx.Tracer.End()
//
//	c.Unset(ctx)
//
//	// 如果配置文件reload 那么对象不放回pool中
//	if ctx.LocalVersion != ctx.Version.Version() {
//		return
//	}
//
//	c.P.Put(ctx)
//}
//
//func (c *CorePool) GetObj() *Core {
//	for {
//		context := c.P.Get().(*Core)
//		if context.LocalVersion != context.Version.Version() {
//			continue
//		}
//		return context
//	}
//}
//
//var registerList []Register

type Register interface {
	With(o ...Option) Register
	Reload(c *Core)
	Set(c *Core)
	Unset(c *Core)
}

type Pooler interface {
	RegisterList(l []Register)
	Set(c *Core)
	Unset(c *Core)
	Reload()
	GetInit()
	ReUse(c *Core)
	GetObj() *Core
}
