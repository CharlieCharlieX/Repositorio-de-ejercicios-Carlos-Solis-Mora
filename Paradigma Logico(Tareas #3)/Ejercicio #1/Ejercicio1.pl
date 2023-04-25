%%Trabajo realizado por Carlos Solis Mora. Carne: 2021051459.
sub_cadenas(Subcadena, Lista, Filtradas) :-
    findall(Cadena, (member(Cadena, Lista), sub_atom(Cadena, _, _, _, Subcadena)),
            Filtradas).

