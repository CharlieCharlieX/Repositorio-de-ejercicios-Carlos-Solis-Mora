// For more information see https://aka.ms/fsharp-console-apps

//Este es el valor mínimo que deben tener los productos para poder aplicárseles el impeusto de venta elegido.
let rangoPagoImpuestos = 2000

//Este es el porcentaje del impuesto, correspondiente.
let porcentajeImpuesto = 0.13

//Esta es la estructura Producto, con sus atributos correspondientes. Teniendo un nombre tipo string, una descripción tipo string
//y un MontoVenta de tipo int.
type Producto = {
    Nombre: string
    Descripcion: string
    MontoVenta: int
}

//Esta es la lista con los productos creados.
let factura = [
    { Nombre = "Llanta"; Descripcion = "Repuestos de autos"; MontoVenta = 1000 }
    { Nombre = "Mouse"; Descripcion = "Perifericos"; MontoVenta = 2000 }
    { Nombre = "Taclado"; Descripcion = "Perifericos"; MontoVenta = 3000 }
    { Nombre = "Motor"; Descripcion = "Repuestos de autos"; MontoVenta = 4000 }
    { Nombre = "Calculadora"; Descripcion = "Útiles escolares"; MontoVenta = 5000 }
    { Nombre = "Cuaderno"; Descripcion = "Útiles escolares"; MontoVenta = 6000 }
]

//Esta es la función encargada de devolver el precio total de una factura. Realiza un map del atributo del MontoVenta
//y luego a la lista resultante se le realiza un ".sum", lo que devuelve el valor total.
let calcularMontoFactura (factura: Producto list) =
    let montos = List.map (fun p -> p.MontoVenta) factura
    List.sum montos

//Esta función es la que se encarga de calcular el precio total de los productos que exceden rangoPagoImpuestos establecido.
//Realiza un filter a los productos que tengan un MontoVenta mayor al rangoPagoImpuestos, y luego se realiza un map del
//MontoVenta multiplicado por el impuesto determinado en la constante porcentaje Impuesto de esos productos,para por último
//realizar a la lista resultante un ".sum" que devuelve el valor total del proceso.
let calcularPrecioImpuesto (factura: Producto list) =
    let productosImpuestos = List.filter (fun p -> p.MontoVenta > rangoPagoImpuestos) factura
    let impuestos = List.map (fun p -> int32(float p.MontoVenta * porcentajeImpuesto)) productosImpuestos
    List.sum impuestos


//Esta es la llamada par a ver el monto total de la factura creada
printfn "El valor final de la factura es: %A" (calcularMontoFactura factura)
//Esta es la llamada para ver el monto total de los productos a los cuales se les puede aplicar el impuesto de venta (También conocido como factura por concepto de impuesto de venta).
printfn "El valor final de la factura por concepto del 13 pociento de impuesto de venta es: %A" (calcularPrecioImpuesto factura)