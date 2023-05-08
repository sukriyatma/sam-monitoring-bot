
$(document).ready( ()=> {
    if (document.location.href.indexOf("monitor.yatma.xyz") === -1 ) return
    localStorage.removeItem("usr"); localStorage.removeItem("pwd")
    $("#formLogin").on("click", "#loginButton", async ()=> {
        var usr = $("#inputUsername").val(); 
        var pwd = $("#inputPassword").val();
        if (usr == "" || pwd === "") {$("#replyForm").text("Username or Password must filled");return}
        $("#replyForm").text("")
        $("#replyForm").css("color", "#FFF")

        axios.get(`http://api.perritosen.com/monitoringbot/login?username=${usr}&password=${pwd}`)
        .then((response) => {
            window.location.href = "home"
            localStorage.setItem("usr", usr)
            localStorage.setItem("pwd", pwd)
        }).catch(error => {
            if (error.code === "ERR_NETWORK") {
                $("#replyForm").text("Check your connection")
                $("#replyForm").css("color", "#FF4646")
                return
            }
            if (error.response.status >= 400 && error.response.status < 500) {
                $("#replyForm").text("Username or Password is wrong")
                return
            }
            if (error.response.status >= 500) {
                $("#replyForm").text("Server Koid, wait admin to fix it")
                return
            }
        })   
    
    })

})

$(document).ready( ()=> {
    if (document.location.href.indexOf("home") === -1 ) return
    if ( !localStorage.getItem("usr") ) {window.location.href = "http://monitor.yatma.xyz/"; return }

    $("#usernameUser").html(localStorage.getItem("usr"))
    $("#logOutButton").click( ()=> {localStorage.removeItem("usr");localStorage.removeItem("pwd"); window.location.href="index.html"})
    $("#ulMonitor").on("click", "#divDataMonitor", (data) => {
        getAllBots(data.currentTarget.parentElement.id)
        setMonitorActive(data.currentTarget.parentElement.id)
        setMonitorName(data.currentTarget.parentElement.id)
    })
    $("#btnTotalOnline").click( () => {
        getBotsByStatus($("#monitorName").html(), "ONLINE")
    })
    $("#btnTotalOffline").click( () => {
        getBotsByStatus($("#monitorName").html(), "OFFLINE")
    })
    $("#btnTotalLoginFailed").click( () => {
        getBotsByStatus($("#monitorName").html(), "LOGIN_FAILED")
    })
    $("#btnTotalBan").click( () => {
        getBotsByStatus($("#monitorName").html(), "BAN")
    })
    $("body").on("click", "img.removeMonitorButton", (data)=> {
        removeMonitor(data.currentTarget.id)
    })

    $("body").on("click", "#nameBot", (data)=> {
        if ($(data.currentTarget).children(".expand-icon").attr("id") === "hidden" ) {
            $($(data.currentTarget.parentElement).children("#detailsBot")).show()
            $($(data.currentTarget.parentElement.parentElement.parentElement)).css("height", "40%")
            $($(data.currentTarget.parentElement.parentElement.parentElement)).css("overflow", "visible")
            $(data.currentTarget).children(".expand-icon").attr("id", () => {
                return "show"
            })
            $(data.currentTarget).children(".expand-icon").css({"transform": "rotate(180deg)", "transition": "transform 0.5s ease"})
            $(data.currentTarget).css("border-radius", "10px 10px 0px 0px")
            $($(data.currentTarget).children("#nameBot span")).css("color", "#FCFF75")
            return
        }

        $($(data.currentTarget.parentElement).children("#detailsBot")).hide()
        $($(data.currentTarget.parentElement.parentElement.parentElement)).css("height", "auto")
        $($(data.currentTarget.parentElement.parentElement.parentElement)).css("overflow", "hidden")
        $(data.currentTarget).children(".expand-icon").attr("id", () => {
            return "hidden"
        })
        $(data.currentTarget).children(".expand-icon").css({ "transform": "rotate(0deg)", "transition": "transform 0.5s ease"})
        $(data.currentTarget).css("border-radius", "10px 10px 10px 10px",)
        $($(data.currentTarget).children("#nameBot span")).css("color", "#FFFFFF")
    })
    

    var getAllMonitors = ()=> {
        axios.get(`http://api.perritosen.com/monitoringbot/findmonitors?username=${localStorage.getItem("usr")}&password=${localStorage.getItem("pwd")}`)
        .then(response => {
            if (response.data === null) return
            setUpMonitor(response.data)
            getAllBots(response.data[1])
            setMonitorActive(response.data[1])
            setMonitorName(response.data[1])
        })
        .catch(console.error)
    }

    var getAllBots = async (monitor) => {
        axios.get(`http://api.perritosen.com/monitoringbot/getbots?username=${localStorage.getItem("usr")}&password=${localStorage.getItem("pwd")}&monitor=${monitor}`)
        .then(setUpBots)
        .catch(console.error)
    }

    var getBotsByStatus = (monitor, status) => {
        axios.get(`http://api.perritosen.com/monitoringbot/findbotsbystatus?username=${localStorage.getItem("usr")}&password=${localStorage.getItem("pwd")}&monitor=${monitor}&status=${status}`)
        .then(setUpBots)
        .catch(console.error)
    }

    var removeMonitor = (monitor) => {
        axios.post(`http://api.perritosen.com/monitoringbot/removemonitor?username=${localStorage.getItem("usr")}&password=${localStorage.getItem("pwd")}&monitor=${monitor}`)
        .then(getAllMonitors)
        .catch(console.error);
    }
    
    var setUpBots = (response) => {
        const totalInfoBot = {"ONLINE":0,"OFFLINE":0,"LOGIN_FAILED":0,"BAN":0}
        let HTMLulListBots = ""
        if (!response.data.list) {
            return
        }

        response.data.list.forEach( element => {
            
            HTMLulListBots +=
                `<li id="liListBot">
                <div id="divBot">
                    <div id="divIndicatorStatus">
                        <span class="indicatorStatus" id="status${element.status}"></span>
                    </div>

                    <div style="width: 100%; height: 100%; display: flex; flex-direction: column; justify-content: center;">
                        <section id="nameBot">
                            <span>${element.name}</span>
                            <img class="expand-icon" id="hidden" src="aset/image/expand.png" style="float: right;">
                        </section>
                        <section id="detailsBot">
                            <div id="divDetailsBot">
                                <table width="90%">
                                    <tr class="headerBot1">
                                        <td width="200px">
                                            <div>
                                                <span class="botDataHeader">Status</span>
                                                <span class="botData" id="statusBot">${element.status}</span>
                                            </div>
                                        </td>
                                        <td width="200px">
                                            <div>
                                                <span class="botDataHeader" >Captcha</span>
                                                <span class="botData" id="captchaBot">${element.captcha}</span>
                                            </div>
                                        </td>
                                        <td width="200px">
                                            <div>
                                                <span class="botDataHeader" >Current World</span>
                                                <span class="botData" id="worldBot">${element.world}</span>
                                            </div>
                                        </td>
                                    </tr>
                                    <tr class="headerBot2">
                                        <td width="200px">
                                            <div>
                                                <span class="botDataHeader" >Position</span>
                                                <span class="botData" id="positionBot">${element.x},${element.y}</span>
                                            </div>
                                        </td>
                                        <td width="200px">
                                            <div>
                                                <span class="botDataHeader" >Level</span>
                                                <span class="botData" id="levelBot">${element.level}</span>
                                            </div>
                                        </td>
                                        <td width="200px">
                                            <div>
                                                <span class="botDataHeader" >Last Updated</span>
                                                <span class="botData" id="lastupdateBot">13:43</span>
                                            </div>
                                        </td>
                                    </tr>
                                </table>
                            </div>
                        </section>
                    </div>    
                </div>
            </li>`
            
            $("#ulListBots").html(HTMLulListBots)
            totalInfoBot[element.status] += 1
        });

        $("section#detailsBot").hide()
        $("li#liListBot").css("height", "auto")
        $("li#liListBot").css("overflow", "hidden")

        if (totalInfoBot.ONLINE > 0) {
            $("#totalOnline").html(totalInfoBot.ONLINE)
            $("#totalOnline").css("color","#9EFF7B")
        }
        if (totalInfoBot.OFFLINE > 0) {
            $("#totalOffline").html(totalInfoBot.OFFLINE)
            $("#totalOffline").css("color","#717171")
        }
        if (totalInfoBot.LOGIN_FAILED > 0) {
            $("#totalLoginFailed").html(totalInfoBot.LOGIN_FAILED)
            $("#totalLoginFailed").css("color","#FFC147")
        }
        if (totalInfoBot.BAN > 0) {
            $("#totalBan").html(totalInfoBot.BAN)
            $("#totalBan").css("color","#FF4646")
        }
    }

    var setUpMonitor = (monitors) => {
        let HTMLulMonitor = "";
        for (let i=0; i<monitors.length; i++) {
            if (monitors[i] === "") continue;
            HTMLulMonitor += 
            `<li class="liDataMonitor" id="${monitors[i]}">
                <div id="divDataMonitor">
                    <span>${monitors[i]}</span>
                </div>
                <img class="removeMonitorButton" id="${monitors[i]}" src="aset/image/trash.png">
            </li>`;
        }
        $("#ulMonitor").html(HTMLulMonitor);
    }

    var setMonitorActive = (monitor) => {
        $(`.liDataMonitor div`).css("background-color", "#363A44");
        $(`#${monitor} div`).css("background-color", "#22242B");
    }

    var setMonitorName = (monitor) => {
        $("#monitorName").html(monitor);
    }

    getAllMonitors();
    setInterval(() => {
        getAllBots($("#monitorName").html());
    }, 30 * 1000);
})


