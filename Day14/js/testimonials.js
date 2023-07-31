// PAKE OOP JAVASCRIPT

// class Testimonials{
//     #quote = ""
//     #image = ""

//     constructor(quote, image){
//         this.#quote = quote
//         this.#image = image
//     }

//     get quote(){
//         return this.#quote
//     }

//     get image(){
//         return this.#image
//     }

//     get author(){
//         throw new Error("Harus Ada Author Untuk membuat Testimonials")
//     }

//     get testimonialHTML(){
//         return `
//         <div class="card-testimonials">
//             <img src="${this.image}"/>
//             <p class="quote">"${this.quote}"</p>
//             <p class="author">- ${this.author}</p>
//         </div>`
//     }


// }

// class AuthorTestimonials extends Testimonials {
//     #user = ""

//     constructor(user, quote, image){
//         super(quote, image)
//         this.#user = user;
//     }

//     get author(){
//         return  this.#user + " (User)"
//     }
// }

// class CompanyTestimonials extends Testimonials {
//     #company = ""

//     constructor(company, quote, image){
//         super(quote, image)
//         this.#company = company;
//     }

//     get author(){
//         return this.#company + " (Company)"
//     }
// }

// const testimonial1 = new AuthorTestimonials("Naruto Uzumaki", "Dattebayo", "https://i.pinimg.com/originals/c0/ff/5b/c0ff5be1deb59a06e0d4a4c303e986cf.jpg")

// const testimonial2 = new AuthorTestimonials("Tanjiro", "Hinokami No Kata", "https://www.dailysia.com/wp-content/uploads/2022/02/Kamado-Tanjiro_1.jpg")

// const testimonial3 = new CompanyTestimonials("Eren Jeager", "TATAKAE", "https://www.dailysia.com/wp-content/uploads/2022/01/Eren-Yeager_1.jpg")

// const testimonial4 = new Testimonials("Kocak", "")


// let dataTestimonial = [testimonial1, testimonial2, testimonial3];

// let testimonialHTML = "";

// for (let i = 0; i < dataTestimonial.length; i++) {
//     testimonialHTML += dataTestimonial[i].testimonialHTML
// }

// document.getElementById("container-testimonials").innerHTML = testimonialHTML;

// 
// HOF AND CALLBACK (foreach terutama)
// 

let testimonialData = [
    {
        user : "Naruto Uzumaki",
        quote : "Ini Sangat Kurang Dattebayo",
        image : "https://i.pinimg.com/originals/c0/ff/5b/c0ff5be1deb59a06e0d4a4c303e986cf.jpg",
        rating : 2
    },
        
    {
        user : "Tanjiro",
        quote : "Lumayan Keren, Ayo Belajar Teknik Pedang",
        image : "https://www.dailysia.com/wp-content/uploads/2022/02/Kamado-Tanjiro_1.jpg",
        rating : 3
    },

    {
        user : "Saitama",
        quote : "Oke Lah",
        image : "https://4kwallpapers.com/images/wallpapers/saitama-one-punch-man-2560x2560-10084.jpg",
        rating : 4
    },

    {
        user : "Eren Jeager",
        quote : "BAGUS BANGET, AYO HANCURKAN DUNIA",
        image : "https://www.dailysia.com/wp-content/uploads/2022/01/Eren-Yeager_1.jpg",
        rating : 5
    },

    {
        user : "Gojo Satoru",
        quote : "Datte Kimi, Yowaimo",
        image : "https://fictionhorizon.com/wp-content/uploads/2023/04/gojo.jpg",
        rating : 5
    }
]

function allTestimonial(){
    let testimonialHTML = "";

    testimonialData.forEach((data) => {
        testimonialHTML += `
        <div class="card-testimonials">
            <img src="${data.image}"/>
            <p class="quote">"${data.quote}"</p>
            <p class="author">- ${data.user} | ${data.rating} <i class="fa-solid fa-star"></i></p>
        </div>`
    })

    document.getElementById("container-testimonials").innerHTML = testimonialHTML

}



allTestimonial();

function filterTestimonialByRate(rating){
    let filteredTestimonialHTML = "";

    const filteredData = testimonialData.filter((data) =>{
        return data.rating === rating
    })

    filteredData.forEach((data) =>{
        filteredTestimonialHTML += `
        <div class="card-testimonials">
            <img src="${data.image}"/>
            <p class="quote">"${data.quote}"</p>
            <p class="author">- ${data.user} | ${data.rating} <i class="fa-solid fa-star"></i></p>
        </div>`
    })

    document.getElementById("container-testimonials").innerHTML = filteredTestimonialHTML
}