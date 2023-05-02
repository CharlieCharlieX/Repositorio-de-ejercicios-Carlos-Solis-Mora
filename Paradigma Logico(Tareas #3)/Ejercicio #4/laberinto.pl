%%Trabajo realizado Por Carlos Eduardo Solis Mora. Carne: 2021051459

%%Laberinto utilizado:
conectado(i,2).
conectado(2,3).
conectado(2,8).
conectado(8,9).
conectado(9,3).
conectado(3,4).
conectado(4,10).
conectado(10,16).
conectado(16,22).
conectado(22,21).
conectado(21,15).
conectado(15,14).
conectado(14,13).
conectado(13,7).
conectado(7,1).
conectado(14,20).
conectado(20,26).
conectado(22,28).
conectado(26,27).
conectado(27,28).
conectado(28,29).
conectado(29,23).
conectado(23,17).
conectado(17,11).
conectado(11,5).
conectado(5,6).
conectado(28,34).
conectado(34,35).
conectado(35,36).
conectado(36,30).
conectado(30,24).
conectado(24,18).
conectado(18,12).
conectado(32,31).
conectado(31,25).
conectado(25,19).
conectado(34,33).
conectado(33,32).
conectado(32,f).

conectados(X,Y):- conectado(X,Y).
conectados(X,Y):- conectado(Y,X).


%! Ini: Nodo origen para la búsqueda
%  Fin: Nodo finalo para la búsqueda
%  LRuta: Lista auxiliar que se irá llenando hacia adelante en recusión
%  hasta encontrar o no solución
%  Res: Variables donde se unificarán los resultados obtenidos.

ruta(Fin,Fin,LRuta,Res):-
    reverse([Fin|LRuta],Res). %% reversa al resultado obtenido
ruta(Ini,Fin,LRuta,Res):-
    conectados(Ini,Z),
    not(member(Z,LRuta)),
    ruta(Z,Fin,[Ini|LRuta],Res).

% Encuentra todos los caminos entre dos nodos dados
% Uso: todos_los_caminos(Inicio, Fin, ListaDeCaminos)
% Ejecutar esto para ver el funcionamiento del código.
todos_los_caminos(Inicio, Fin, ListaDeCaminos) :-
    findall(Camino, ruta(Inicio, Fin, [], Camino), ListaDeCaminos).


