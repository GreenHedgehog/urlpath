# urlpath

## Supported types
Now, only primitives types are supported(`string`, `int*`, `uint*`, `bool`). But, it could be easily extended with `Marshaler`/`Unmarshaler` realization for custom types. Types `Base64` and `URLEscaped` are predefined already.

```go
type _ struct {
    _ int
    _ int8
    _ int16
    _ int32
    _ int64

    _ uint
    _ uint8
    _ uint16
    _ uint32
    _ uint64

    _ string
    _ bool

    _ urlpath.Base64
    _ urlpath.URLEscaped
}
```

## Tags

```go
type _ struct {
    Field Type `urlpath:""`                    // Fields without defined package tags would be ignored for Marshaling/Unmarshaling
    Field Type `urlpath:"-"`                   // Expilicit ignoring field for Marshaling/Unarshaling
    Field Type `urlpath:"name"`                // Defining urlpath name for field
    Field Type `urlpath:"name;required"`       // Defines field as required for appearing in parsed urlpath, otherwise would rise an error
    Field Type `urlpath:"name;omitempty"`      // If value of field is a zero-value, then it would be omitted in urlpath (default, it would appear with zero value)
    Field Type `urlpath:"name;default=12345"`  // If value of field is a zero-value, then it would be replaced with that default value during Marshaling/Unmarshaling
    Field Type `urlpath:"name;scheme=one,two"` // Scheme tag allows to use different strategy of marshaling
}
```

Example of scheme usage:

```go

type Products struct {
    Apple int     `urlpath:"apple;scheme=fruit,red"`
    Pear int      `urlpath:"pear;scheme=fruit,green"`
    Raspberry int `urlpath:"raspberry;scheme=berry,red"`
}

p := Products{1, 2, 3}
  
_ = urlpath.MarshalScheme("fruit", p) // /apple/1/pear/2  
_ = urlpath.MarshalScheme("red", p) // /apple/1/raspberry/3

```

## Limitations

Fields with type base on `string` always would be with `omitempty` flag (implicit), because, otherwise, it leads to double slashes in urlpath, which could rise some issues. 

