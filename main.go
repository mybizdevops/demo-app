package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("demo").Parse(`<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>myBIZDEVOPS</title>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <style>
@import url('https://fonts.googleapis.com/css?family=Arvo');

:root {
    --color-neutral: rgb(76, 76, 76);
    --color-biz: rgb(11, 177, 55);
    --color-ops: rgb(236, 55, 55);
    --color-dev: rgb(66, 134, 244);

    --grid-gap: 15px;
}

html, body {
    font-family: 'Arvo', serif;
    color: var(--color-neutral);
    font-size: 18px;
    background-color: rgba(254,254,254,1);
    line-height: 28px;
}

* {
    margin: 0;
    padding: 0;
}

.header {
    width: 100%;
    height: 80px;
    position: fixed;
    font-size: 26px;
    background-color: rgba(254,254,254,0.8);
}

nav {
    padding-left: 20px;
    padding-right: 20px;
    padding-top: 24px;
    margin: auto;
    max-width: 800px;
}

a {
    color: var(--color-neutral);
}

.logo {
    float: left;
}

.github {
    float: right;
}

.logo>.biz, .logo>.dev, .logo>.ops {
    letter-spacing: 0.2em;
    font-weight: bold;
}

.content {
    padding-left: 20px;
    padding-right: 20px;
    padding-top: 100px;
    margin: auto;
    max-width: 800px;
}

.splashimg {
    padding-top: 100px;
    padding-bottom: 100px;
    width: 70%;
    margin: auto;
}

.biz {
    color: var(--color-biz);
}

.bizelems > path {
    fill: var(--color-biz);
}

.dev {
    color: var(--color-dev);
}

.develems > path {
    fill: var(--color-dev);
}

.ops {
    color: var(--color-ops);
}

.opselems > path {
    fill: var(--color-ops);
}

.neutral {
    color: var(--color-neutral);
}
        </style>
    </head>
    <body>
        <div class="main">
            <div class="header">
                <nav>
                    <div class="logo">
                        <span class="neutral">MY</span><span class="biz">BIZ</span><span class="dev">DEV</span><span class="ops">OPS</span>
                    </div>
                    <div class="github"><a href="https://www.github.com/mybizdevops">GitHub</a></div>
                </nav>
            </div>

            <div class="content">
                <p>Given you can see this the demo was successful! This website is running on a server with the hostname <b>{{.}}</b>.</p>
                <div class="splashimg">
                    <svg id="bizdevops" version="1.1" viewBox="0 0 35.919052 11.50958">
                        <g transform="translate(-53.143429,-106.12292)" id="layer1">
                        <g class="develems">
                        <path d="m 74.162424,106.29829 -1.595005,0.0991 0.0991,1.59501 1.595004,-0.0991 z m -3.897735,0.24216 -1.595004,0.0991 0.0991,1.59501 1.595004,-0.0991 z" />
                        <path d="m 72.173336,109.42497 -1.183049,0.0735 0.225105,3.62309 1.183049,-0.0735 z" />
                        <path d="m 72.113896,116.4896 c 2.01752,-0.12535 3.624653,-0.88263 5.003588,-2.24076 l -0.606426,-0.71519 c -1.380189,1.16734 -2.532393,1.7373 -4.465415,1.8574 -1.933023,0.1201 -3.146927,-0.30285 -4.660987,-1.29037 l -0.513245,0.78476 c 1.536443,1.17699 3.224956,1.72951 5.242485,1.60416 z" />
                        </g>
                        <g class="bizelems">
                        <path d="m 60.359499,106.12292 -1.59063,0.15413 0.154135,1.59064 1.590629,-0.15414 z m -3.732909,1.96729 -0.154134,-1.59063 -1.59063,0.15413 0.07656,0.79005 c -0.558305,0.0541 -0.908666,-0.0502 -1.239706,-0.61354 l -0.575251,0.42789 c 0.479591,0.88923 1.102485,1.06279 2.018942,0.97399 z" />
                        <path d="m 59.241485,111.80066 -1.632325,1.04071 0.140864,1.45369 3.177302,-2.10484 -0.101055,-1.04287 -3.523338,-1.46618 0.135761,1.40102 z" />
                        <path d="m 53.407478,114.17969 0.09391,0.96913 8.055856,2.48368 -0.09391,-0.96913 z" />
                        </g>
                        <g class="opselems" transform="translate(-2.1166667,0.52916667)">
                        <path d="m 90.867526,106.14247 -1.595541,-0.0901 -0.09013,1.59554 1.59554,0.0901 z m -3.899034,-0.22025 -1.59554,-0.0901 -0.09013,1.59554 1.595539,0.0901 z" />
                        <path d="m 88.51909,109.07468 -1.183446,-0.0668 -0.204733,3.6243 1.183446,0.0669 z" />
                        <path d="m 87.585139,116.85599 3.59401,-3.19964 -0.697872,-0.78144 -2.824931,2.53287 -2.532278,-2.83549 -0.759706,0.6885 z" />
                        </g>
                        </g>
                    </svg>
                </div>
            </div>

            <div class="footer">
            </div>
        </div>
    </body>
</html>`))
	hostname, _ := os.Hostname()
	tmpl.ExecuteTemplate(w, "demo", hostname)
}

func main() {
	http.HandleFunc("/", serveTemplate)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
