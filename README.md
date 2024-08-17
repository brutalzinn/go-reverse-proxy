# go-reverse-proxy

# IN DEVELOPMENT 

### When horses comes with power to disrupt the time machine just for play minecraft with mods and all money is spend with random eletronics products.

#### This sentence doesnt have any sense. This project too.

##### The glorisus proposite of this project is to handle like a router between TCP, UDP and HTTP protocols. Why? If you uses a proxy server like ngrok or localtonet, you pay for every port shared and every channel created. This project joins some workarounds to bypass my money disrupture with idiot things like me searching about servers used to spend money with more idiot things. This just save 0,34 dollar cents with Localtonet today. 


![Loki](https://i.giphy.com/media/v1.Y2lkPTc5MGI3NjExZmU5YWFxdG03NXFkM3ljdTR6cmloOXg4cnN0eDM2NGFvMjlkczE1YiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/VlrRiZAur3jylN2XKT/giphy.gif)



###### Disclaimer 1: if this readme is soo cool, are you agree with me that this probabily dont will work as expected and not be tested properly for multiplataform assigment? BUT we know this works at my machine and you need to agree with me.


###### Disclaimer 2: Its just a project that i started at 1:30 AM and now is 08:50 AM and like every war we lost, this is the trully glory ( now my friends are sleeping and no one played minecraft again)


# OK. I undestand that everything here is garbage and lost of time.

YEEP. BUT keep in mind that this is just the 1ª project of this week.  A series of projects that i will develop that is useless BUT just only for now.


# How to setup this sh#@t?!


1. I am assuming you has docker and make at you machine ( if you dont uses these tools, you are shrek here.)

![Shrek](https://media1.tenor.com/m/mtiOW6O-k8YAAAAd/shrek-shrek-rizz.gif)


1. make build

2. make run

3. Edit the config.json generated with you especification:

        {
            "geral":{
                "port":"8080"
            },
            "routes":[
                {
                    "IN":{
                        "host":"127.0.0.1",
                        "port":"25565",
                        "protocol":"tcp",
                    },
                    "OUT":{
                        "host":"127.0.0.1",
                        "port":"25568",
                        "protocol":"tcp"
                    }
                }
            ]
        }
        /// this case every access by TCP at port 25565 will be redirect to 25568.



# Me and my cat at 3:55 AM wondering how we got to this

![Loki and Silvy](https://www.slashfilm.com/img/gallery/kevin-feige-confirms-loki-and-sylvie-broke-the-multiverse-thanks-a-lot-guys/intro-1651609635.webp)