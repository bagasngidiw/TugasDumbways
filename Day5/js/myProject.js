function whenFormEmpty(){
    let name = document.getElementById("input-project-name").value;
    let startDate = document.getElementById("input-start-date").value;
    let endDate = document.getElementById("input-end-date").value;
    let description = document.getElementById("input-description").value;
    let multiCheckbox = document.querySelectorAll(".input-technologies:checked");
    let image = document.getElementById("input-image").value;

    console.log(name)
    console.log(startDate)
    console.log(endDate)
    console.log(description)
    console.log(multiCheckbox)
    console.log(image)



    // validation user form
    if(name === ""){
        return alert("Please Fill your Project Name");
    }else if (startDate === ""){
        return alert("Please Fill When your Project Start");
    }else if(endDate === ""){
        return alert("Please Fill When your Project Finish");
    }else if(description === ""){
        return alert("Please Write your Project Description");
    }else if (multiCheckbox.length === 0) {
        return alert("Please Choose your Project Technologies");
    }else if(image === ""){
        return alert("Plese Attach your Project Image");
    }

}

let projectData = []
let defaultProjectData = document.getElementById("list-project-container").innerHTML;

function addProjectData(event){

    event.preventDefault();

    let name = document.getElementById("input-project-name").value;
    let startDate = document.getElementById("input-start-date").value;
    let endDate = document.getElementById("input-end-date").value;
    let description = document.getElementById("input-description").value;
    let image = document.getElementById("input-image").files;

    
    const nodejsIcon = `<i class="fa-brands fa-node-js fa-xl tooltip"><span class="tooltiptext tooltip-bottom">Node Js</span></i>`;
    const reactjsIcon = `<i class="fa-brands fa-react fa-xl tooltip"><span class="tooltiptext tooltip-bottom">React Js</span></i>`;
    const pythonIcon = '<i class="fa-brands fa-python fa-xl tooltip"><span class="tooltiptext tooltip-bottom">Pyhton</span></i>';
    const javaScriptIcon = '<i class="fa-brands fa-square-js fa-xl tooltip"><span class="tooltiptext tooltip-bottom">JavaScript</span></i>';

    const nodeJs = document.getElementById("nodejs");
    const reactJs = document.getElementById("reactjs");
    const python = document.getElementById("python");
    const javaScript = document.getElementById("javaScript");

    // IBARATNYA SI ICON SET INI NYIAPIN 1 KOTAK BUAT PARA ICON DIATAS KETIKA DI CEKLIS
    let iconSet = "";

    // VALIDASI IKON KETIKA DI CEK  
    if(nodeJs.checked){
        iconSet += nodejsIcon;
    }
     if(reactJs.checked){
        iconSet += reactjsIcon;
    } 
    if(python.checked){
        iconSet += pythonIcon
    } 
    if(javaScript.checked){
        iconSet += javaScriptIcon
    }

    // untuk membuat object file menjadi URL secara sementara, agar tampil
    image = URL.createObjectURL(image[0])

    let start = new Date(startDate);
    let end = new Date(endDate);

    if(start > end){
        return alert("Please Put Date Correctly");
    }

    let timeDifference = end.getTime() - start.getTime();
    let days = timeDifference / (1000 * 60 * 60 * 24);
    let weeks = Math.floor ( days/7 );
    let months = Math.floor (weeks/4);
    let years = Math.floor (months/12);

    let duration = "";

    // validation time data
    if(days > 0){
        duration = `${days} Day`
    }
    if(weeks > 0){
        duration = `${weeks} Week`
    }
    if(months > 0){
        duration = `${months} Month`
    }
    if(years > 0 ){
        duration = `${years} Year`
    }
    

    let data = {
        name,
        duration,
        iconSet,
        description,
        image
    };

    projectData.push(data);
    console.log(projectData);
    
    renderProjectCard();

    let a = document.createElement('a');
    a.href = `#list-project-container`;
    a.click();
}

function renderProjectCard(){
    let projectContainer = document.getElementById("list-project-container");
    let currentProjectHtml = defaultProjectData;
    for(let index = 0; index < projectData.length; index++){
        currentProjectHtml += `
                <div class="list-project-card">
                    <div class="card-image">
                        <img src="${projectData[index].image}"/>
                    </div>
                    <div class="card-title">
                        <h4><a href= "projectDetail.html" target ="_blank">${projectData[index].name}</a></h4>
                    </div>
                    <div class="card-duration">
                        <p>Duration : ${projectData[index].duration}</p>
                    </div>
                    <div class="card-description">
                        <p>
                        ${projectData[index].description}
                        </p>
                    </div>
                    <div class="card-icons">
                        ${projectData[index].iconSet}
                    </div>
                    <div class="card-button">
                        <button>Edit</button>
                        <button>Delete</button>
                    </div>
                </div>`
    }

    projectContainer.innerHTML = currentProjectHtml;
}