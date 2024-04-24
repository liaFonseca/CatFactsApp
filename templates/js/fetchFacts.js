
const btnGetFacts = document.getElementById('get-facts-btn');
console.log('entrou aqui')
const getFacts = () => {
    const count = document.getElementById('count').valueAsNumber 
    const lang = document.getElementById('lang').value

    axios.get('/api/catFacts', {
        params: {
            count : count,
            lang : lang
        }
    }).then((response) => {
        if (response.status === 200){
            location.reload()
        }
    }) 
};

btnGetFacts.addEventListener('click', getFacts)