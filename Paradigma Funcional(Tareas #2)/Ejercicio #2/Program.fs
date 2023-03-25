// For more information see https://aka.ms/fsharp-console-apps

(*
Trabajo realizado por Carlos Eduardo Solís Mora.
Carné: 2021051459
*)

//Esta función se encarga de generar una lista, creando en la primera posición el primer número que se pasa por parámetro
//Y la última posición corresponde al segundo numero pasado por parámetro. Luego se llama de manera recursiva las posiciones
//intermedias sumandole uno al primer valor hasta que se complete la lista.
let rec general_seq a b =
    if a > b then []
    else a :: general_seq (a + 1) b
    
let seq1 = general_seq 1 5
printf "La lista es %A" seq1
