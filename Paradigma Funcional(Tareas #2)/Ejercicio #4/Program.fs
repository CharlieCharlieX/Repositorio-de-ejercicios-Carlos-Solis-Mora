// For more information see https://aka.ms/fsharp-console-apps

(*
Trabajo realizado por Carlos Eduardo Solís Mora.
Carné: 2021051459
*)

// Ejercicios de rutas en Grafos con búsqueda en profundidad con peso

//        a ---- c ---- x
//      /   \  /
//     /     \/
//   i       X
//     \    / \
//      \  /   \
//        b ---- d ---- f
let grafo = [
    [('i', 0);('a', 1);('b', 2)];
    [('a', 1);('i', 0);('c', 3);('d', 2)];
    [('b', 2);('i', 0);('c', 1);('d', 4)];
    [('c', 3);('a', 1);('b', 1);('x', 2)];
    [('d', 2);('a', 4);('b', 1);('f', 3)];
    [('f', 3);('d', 2)];
    [('x', 2);('c', 3)]
]

let miembro e (lista: ('a * 'b) list) =
    lista
    |> List.map (fun (x,_) -> x = e)
    |> List.reduce (fun x y -> x || y)

let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | ((head, peso)::adyacencias)::tail when nodo = head -> adyacencias
    | _::tail -> vecinos nodo tail 

let extender (ruta: _ list) grafo =
    (vecinos (fst ruta.Head) grafo)
    |> List.map (fun (x,peso) -> if (miembro x ruta) then [] else (x,peso)::ruta )
    |> List.filter (fun x -> x <> []) 

(*Esta es la primera función agregada para resolver el problema de encontrar la ruta más corta. Esta función mapea todos
los pesos de las tuplas una ruta en específico y luego suma sus valores con List.sum, para así obtener el peso total de esa
ruta en específico.
*)
let sumaLista (ruta: _ list) grafo =
    ruta
    |> List.map (fun (_,peso) -> peso)
    |> List.sum

(*Esta es la segunda función agregada para resolver el problema de encontrar la ruta más corta. Esta función obtiene
cada una de las rutas almacenadas en una lista y les aplica la función .minBy, utilizando como criterio la función anteriormente
creada "sumaLista", entonces minBy aplicará a cada una de las rutas contenidas la función sumaLista, y devolverá la ruta que
obtuvo el valor más pequeño, logrando así el objetivo de obtener la ruta más corta.
*)
let menor (rutas: _ list list) grafo =
    rutas
    |> List.minBy(fun x -> sumaLista x grafo)

let rec prof_aux fin (rutas: _ list list) grafo =
    if rutas = [] then
        []
    elif fin = (fst rutas.Head.Head) then
        //List.rev rutas.Head
        List.append
            ([List.rev rutas.Head])
            (prof_aux fin rutas.Tail grafo)
    else
        prof_aux fin (List.append (extender rutas.Head grafo) rutas.Tail) grafo
              
        
let prof ini fin grafo =
    prof_aux fin [[(ini,0)]] grafo
    
printfn "La siguiente lista son todas las rutas posibles para llegar de i a f:"
printfn "%A" (prof 'i' 'f' grafo) //Llamada para encontrar en profundidad la lista de rutas que nos llevan de i a f
printfn " "
printfn "Esta es la ruta más corta:"
printfn "%A" (menor(prof 'i' 'f' grafo) grafo) //Esta es la llamada a nuestra función para encontrar la ruta más corta,
                                               //utiliza la lista de rutas para llegar de i a f que se obtiene de la búsqueda
                                               //en profundidad
