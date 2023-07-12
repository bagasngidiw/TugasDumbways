class Testimonials{
    #quote = ""
    #image = ""

    constructor(quote, image){
        this.#quote = quote
        this.#image = image
    }

    get quote(){
        return this.#quote
    }

    get image(){
        return this.#image
    }

    get author(){
        throw new Error("Harus Ada Author Untuk membuat Testimonials")
    }

    get testimonialHTML(){
        return `
        <div class="card-testimonials">
            <img src="${this.image}"/>
            <p class="quote">"${this.quote}"</p>
            <p class="author">- ${this.author}</p>
        </div>`
    }


}



class AuthorTestimonials extends Testimonials {
    #user = ""

    constructor(user, quote, image){
        super(quote, image)
        this.#user = user;
    }

    get author(){
        return  this.#user + " (User)"
    }
}

class CompanyTestimonials extends Testimonials {
    #company = ""

    constructor(company, quote, image){
        super(quote, image)
        this.#company = company;
    }

    get author(){
        return this.#company + " (Company)"
    }
}

const testimonial1 = new AuthorTestimonials("Naruto Uzumaki", "Dattebayo", "https://i.pinimg.com/originals/c0/ff/5b/c0ff5be1deb59a06e0d4a4c303e986cf.jpg")

const testimonial2 = new AuthorTestimonials("Tanjiro", "Hinokami No Kata", "https://www.dailysia.com/wp-content/uploads/2022/02/Kamado-Tanjiro_1.jpg")

const testimonial3 = new CompanyTestimonials("Eren Jeager", "TATAKAE", "https://static.wikia.nocookie.net/shingekinokyojin/images/a/a1/Eren_Jaeger_%28Anime%29_character_image.png/revision/latest?cb=20220123225500")

const testimonial4 = new Testimonials("Kocak", "")


let dataTestimonial = [testimonial1, testimonial2, testimonial3];

let testimonialHTML = "";

for (let i = 0; i < dataTestimonial.length; i++) {
    testimonialHTML += dataTestimonial[i].testimonialHTML
}

document.getElementById("container-testimonials").innerHTML = testimonialHTML;
