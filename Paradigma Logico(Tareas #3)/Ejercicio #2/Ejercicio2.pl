%%Trabajo Realizado por Carlos Solis Mora. Carne: 2021051459.
sub_conjunto([], _).
sub_conjunto(Subconjunto, Conjunto) :-
    subset(Subconjunto, Conjunto).

