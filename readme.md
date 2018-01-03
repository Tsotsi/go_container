## info
> inspired by pimple


## usage

```golang

type Abc struct{
    Cc int
}

c := NewContainer()
c.Set("test", func(cc *Container){
    a := new(Abc)
    a.Cc = 1988
    return a
})

if f,ok:=c.Get("test");ok{
    println(f.(*Abc).Cc == 1988)  // true
    f.(*Abc).Cc += 10
}

if f,ok:=c.Get("test");ok{
    println(f.(*Abc).Cc == 1998)  // true
}


c1 := NewContainer()
c1.Factory("test", func(cc *Container){
    a := new(Abc)
    a.Cc = 1988
    return a
})

if f,ok:=c1.Get("test");ok{
    println(f.(*Abc).Cc == 1988)  // true
    f.(*Abc).Cc += 10
}

if f,ok:=c1.Get("test");ok{
    println(f.(*Abc).Cc == 1988)  // true
}

```