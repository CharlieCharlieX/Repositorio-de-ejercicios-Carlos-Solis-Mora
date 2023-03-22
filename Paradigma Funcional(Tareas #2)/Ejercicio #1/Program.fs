// For more information see https://aka.ms/fsharp-console-apps

open System.Text.RegularExpressions

//La función filtrar palabras utiliza las funcionalidades de la librería Text, en específico RegularExpressions.
//Se utiliza la funcionalidad Regex, que permite darle un formato a una cadena de texto, y entonces utilizamos "\b" que corresponde
//a límites de palabra, lo que significaría que la expresión "\b%s\b" tiene que ser una palabra completa y no simplemente algo
//que coincida bajo cualquier contexto. Para conseguir realizar este formato e imprimirlo en el proceso, utilizamos sprintf
//Para posicionar la palabra en el puntero "%s" que determinamos. Ya por último, vamos a montarlo en una función, que se
//ejecutará cada vez que la función filter determine que hubo una coincidencia del tipo que realizamos, y mantiene el
//elemento en la lista si coincide, y si no, lo descarta.
let filtrarPalabras (pal:string) (lista:string list) =
    let regex = new Regex(sprintf @"\b%s\b" pal)
    lista |> List.filter (fun str -> regex.IsMatch(str))
                    
let frases = ["hola"; "a todos"; "la mama"; "ala"; "la prueba"]
printfn "%A" (filtrarPalabras "la" frases)                   