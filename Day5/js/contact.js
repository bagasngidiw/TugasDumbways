function submitData(){
    let input_name = document.getElementById("input-name");
    let name = input_name.value;

    let input_email = document.getElementById("input-email");
    let email = input_email.value;

    let input_phone = document.getElementById("input-phone");
    let phone = input_phone.value;    

    let input_subject = document.getElementById("input-subject");
    let subject = input_subject.value;

    let input_message = document.getElementById("input-message");
    let message = input_message.value;

    let isError = false;

    if(name === ""){
        input_name.classList.add("error");
        isError = true;
    }else{
        input_name.classList.remove("error");

    }

    if(!email.includes("@")){
        input_email.classList.add("error");
        isError = true;
    }else{
        input_email.classList.remove("error");
    }

    if(phone === ""){
        input_phone.classList.add("error");
        isError = true;
    }else{
        input_phone.classList.remove("error");

    }

    if(subject === ""){
        input_subject.classList.add("error");
        isError = true;
    }else{
        input_subject.classList.remove("error");

    }

    if(message === ""){
        input_message.classList.add("error")
        isError = true;
    }else{
        input_message.classList.remove("error")

    }

    let emailreceiver = "bagasvanbacdim@gmail.com";

    let a = document.createElement('a');
    a.href = `mailto:${emailreceiver}?subject=${subject}&body=Halo Nama Saya ${name}, \n${message}, silahkan kontak saya di nomor berikut : ${phone}`;
    if(!isError){
        a.click()
    }

}

  