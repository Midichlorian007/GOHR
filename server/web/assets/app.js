console.log("Connected")

$("#floatingSelect").change(function () {

    if (this.value === "Golang")
        $("#minLvl").text("Middle"),
            $("#maxLvl").text("Senior"),
            $("#progresPercent").text("80%").width("80%");
    else if (this.value === "Python")
        $("#minLvl").text("Junior"),
            $("#maxLvl").text("Middle"),
            $("#progresPercent").text("90%").width("90%");
    else if (this.value === "Java")
        $("#minLvl").text("Junior"),
            $("#maxLvl").text("Middle"),
            $("#progresPercent").text("10%").width("10%");
    else
        $("#minLvl").text("Junior"),
            $("#maxLvl").text("Middle"),
            $("#progresPercent").text("50%").width("50%");


});

$("#btnFindJob").click(function(){
    $(this).text($(this).text() == 'Уже работаю' ? 'В поиске работы' : 'Уже работаю' );
    $("#jobPlace").toggleClass('d-none')
    $("#jobFunc").toggleClass('d-none')
});


