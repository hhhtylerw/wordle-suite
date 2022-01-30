gameState = item = JSON.parse(localStorage.getItem("gameState"));

hintCount = 0;
async function hint() {
    hintCount++;
    console.log("Hint Count: " + hintCount);
    console.log(gameState);

    switch (hintCount) {
        case 1: // How many vowels in solution
            await hint1();
            break;
        case 2: // One letter in solution
            randomLetter = gameState.solution.charAt(Math.floor(Math.random() * (gameState.solution.length)));
            console.log(randomLetter);
            document.getElementById("hint2").innerHTML = "Hint 2: The solution contains the letter \"" + randomLetter + "\".";
            break;
        case 3: // One letter in solution in correct position
            randomIndex = Math.floor(Math.random() * (gameState.solution.length));
            randomLetter = gameState.solution.charAt(randomIndex);
            document.getElementById("hint3").innerHTML = "Hint 3: The letter \"" + randomLetter + "\" is at index " + (randomIndex + 1) + " in the solution. (1-5)";
            break;
        case 4: // Definition of solution
            await hint4();
            break;
        case 5: // Reveal solution
            document.getElementById("hint5").innerHTML = "Hint 5: The solution is \"" + gameState.solution + ".\"";
            break;
        case 6: // Easter egg
            document.getElementById("hint6").innerHTML = "You already have the solution!";
            break;
    }
}

async function hint1() { // Calculate # of vowels in solution
    vowels = "aeiou";
    vowelsCount = 0;
    for (i = 0; i < gameState.solution.length; i++) {
        if (vowels.includes(gameState.solution.charAt(i))) {
            vowelsCount++;
        }
    }
    document.getElementById("hint1").innerHTML = "Hint 1: There are " + vowelsCount + " vowels in the solution.";
}

async function hint4() { // Read dictionary API 
    console.log("Hint 4");
    document.getElementById("hint4").innerHTML = "Hint 4: Fetching definition...";
    let url = 'https://api.dictionaryapi.dev/api/v2/entries/en/' + gameState.solution;
    fetch(url)
    .then(res => res.json())
    .then(out => {
        console.log('Checkout this JSON! ', out[0].meanings[0].definitions[0].definition);
        document.getElementById("hint4").innerHTML = "Hint 4: The definition of the solution is \"" + out[0].meanings[0].definitions[0].definition + "\"";
    })
    .catch(err => {
        console.log('Error: ', err);
        document.getElementById("hint4").innerHTML = "Hint 4: Error fetching definition.";
    });
}



