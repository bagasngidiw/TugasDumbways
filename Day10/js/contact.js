function submitData(event){

    event.preventDefault();

    let name = document.getElementById("input-name").value;
    // let name = input_name.value;

    let email = document.getElementById("input-email").value;
    // let email = input_email.value;

    let phone = document.getElementById("input-phone").value;
    // let phone = input_phone.value;    

    let subject = document.getElementById("input-subject").value;
    // let subject = input_subject.value;

    let message = document.getElementById("input-message").value;
    // let message = input_message.value;

    // let isError = false;

    if(name === ""){
        return alert("Please Fill your Name");
    }else if(!email.includes("@")){
        return alert("Please Include @ in your Email");

    }else if(phone === ""){
        return alert("Please Fill your Phone Number");

    }else if(subject === ""){
        return alert("Please Select your Skills");

    }else if(message === ""){
        return alert("Please Fill your Message");

    }

    let emailreceiver = "bagasvanbacdim@gmail.com";

    let a = document.createElement('a');
    a.href = `mailto:${emailreceiver}?subject=${subject}&body=Halo Nama Saya ${name}, \n${message}, silahkan kontak saya di nomor berikut : ${phone}`;
    if(!isError){
        a.click()
    }

}

  