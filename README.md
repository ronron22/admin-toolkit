# admin toolkit

```bash
Files
.
|-- chaussette.sh 
|-- crocodile.sh
|-- calumet.sh
`-- README.md
```

## chaussette.sh

*chaussette* interroge les sockets afin de nous donner un top des connexions ouvertes sur un port donné. 

Exemple

Obtenir le top 5 des connexions sur le port 443

```bash
~# bash  chaussette.sh 443 5
63.13.42.244
69.12.124.233
207.6.13.94
173.9.251.160
212.5.11.15
```

Le nombre d'occurence peut être obtenu en ajoutant *stats* à la fin, exemple : `chaussette.sh 443 5 stats` 

## crocodile.sh 

*crocodile* active temporairement une règle firewall limitant l'accès à un port, la limite est à 5 connexions par seconde et s'applique sur toutes les ip sources accèdant à ce port.

deux arguments doivent lui être donné, le port et une durée en minute.

Il s'utilise de la manière suivante :

```bash
~# bash crocodile.sh 443 5
add the "iptables -A INPUT -p tcp --syn --dport 443 -m connlimit \
--connlimit-above 5 -j DROP" rule for 5 minutes
```

Dans l'exemple ci-dessus, *crocodile* limite durant 5 minutes l'accès au port 443. 
 
### voir les jobs en attentes

```bash
atq 
```

### voir le détail d'un job 

```bash
at -c 12
```

## calumet.sh

*calumet* interroge *apache* afin d'obtenir un top des ip connectées et des url accédées.
Le top est de 5 par défaut et est surchargeable.

Exemple

Obtenir le top 5 

```bash
~# bash calumet.sh 5
top 5 of detected ip
---------------------
      7 77.157.16.2 www.bals.com:443 GET
     10 5.23.13.140 www.bals.com:443 GET
     13 90.12.16.48 www.bals.com:443 GET
     20 5.187.5.112 www.bals.com:443 GET
     87 82.64.7.123 www.bals.com:443 GET


top 5 of url called
--------------------
     11 /modal_forms/nojs/login?destination=node/2662%3Ftitle%3D%26 www.bals.com:443
     18 /modal_forms/nojs/login?destination=node/2734%3Ftitle%3D%26 www.bals.com:443
     19 /sites/default/files/styles/color--thumb/public/media/image www.bals.com:443
     21 /modal_forms/nojs/login?destination=node/4722%3Ftitle%3D%26 www.bals.com:443
     64 /sites/default/files/styles/paragraph-multimedia-main/publi www.bals.com:443
```

## combiner les outils

```bash
for i in $(bash chaussette.sh 443 5) ; do
   iptables -I INPUT -s $i -j DROP    
   iptables -D INPUT -s $i -j DROP | at now+1hours
done
```

je bloque les 5 ip les plus représentés durant une heure.
