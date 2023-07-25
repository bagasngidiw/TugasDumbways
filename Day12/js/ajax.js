const promise = new Promise((resolve, reject) => {
    const request = new XMLHttpRequest()

    request.open("GET", "https://api.npoint.io/ff55f026c28de7dac1e2", true)
    request.onload = function(){

        // console.log(request.responseText)

        if(request.status === 200){
            resolve(JSON.parse(request.responseText))
        }else if(request.status >= 400){
            reject("Error Mas Bro")
        }
    }
    request.onerror = function(){
        reject("Jaringan Eror")
    }
    request.send()
})

async function allTestimonial(){
    const dataTestimonial = await promise

    let testimonialHTML = "";
    dataTestimonial.forEach((data) => {
        testimonialHTML += `
        <div class="card-testimonials">
            <img src="${data.image}"/>
            <p class="quote">"${data.quote}"</p>
            <p class="author">- ${data.user} | ${data.rating} <i class="fa-solid fa-star"></i></p>
        </div>`
    })

    document.getElementById("container-testimonials").innerHTML = testimonialHTML

}

// Menampilkan di awal
allTestimonial()

async function filterTestimonialByRate(rating){
    const dataTestimonial = await promise;

    let filteredTestimonialHTML = "";

    const filteredData = dataTestimonial.filter((data) =>{
        return data.rating === rating
    })

    if(filteredData.length === 0){
        filteredTestimonialHTML += `<h1 class="data-not-exist fw-bold">Data is not Exist</h1>`

    }else{
        filteredData.forEach((data) =>{
            filteredTestimonialHTML += `
            <div class="card-testimonials">
                <img src="${data.image}"/>
                <p class="quote">"${data.quote}"</p>
                <p class="author">- ${data.user} | ${data.rating} <i class="fa-solid fa-star"></i></p>
            </div>`
        })
    }

    document.getElementById("container-testimonials").innerHTML = filteredTestimonialHTML
}



