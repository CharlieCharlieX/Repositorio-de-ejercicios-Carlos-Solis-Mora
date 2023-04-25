%%Trabajo realizado Por Carlos Eduardo Solis Mora. Carne: 2021051459
aplanar([], []).
aplanar([X|Xs], Zs) :-
    is_list(X), !,
    aplanar(X, Y), aplanar(Xs, Ys),
    append(Y, Ys, Zs).
aplanar([X|Xs], [X|Ys]) :-
    aplanar(Xs, Ys).
