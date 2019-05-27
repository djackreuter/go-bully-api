DROP TABLE IF EXISTS classes;
DROP TABLE IF EXISTS students;
DROP TABLE IF EXISTS cliques;
DROP TABLE IF EXISTS locations;
DROP TABLE IF EXISTS missions;

CREATE TABLE missions (
    id      INT UNSIGNED AUTO_INCREMENT NOT NULL,
    name    VARCHAR(50)                 NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE locations (
    id      INT UNSIGNED AUTO_INCREMENT NOT NULL,
    name    VARCHAR(50)                 NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE cliques (
    id      INT UNSIGNED AUTO_INCREMENT NOT NULL,
    name    VARCHAR(50)                 NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE students (
    id          INT UNSIGNED AUTO_INCREMENT NOT NULL,
    name        VARCHAR(50)                 NOT NULL,
    clique_id   INT UNSIGNED                NOT NULL,
    FOREIGN KEY (clique_id) REFERENCES cliques (id),
    PRIMARY KEY (id)
);

CREATE TABLE classes (
    id      INT UNSIGNED AUTO_INCREMENT NOT NULL,
    name    VARCHAR(50)                 NOT NULL,
    PRIMARY KEY (id)
);


INSERT INTO missions(name) VALUES ("Welcome to Bullworth");
INSERT INTO missions(name) VALUES ("This Is Your School");
INSERT INTO missions(name) VALUES ("The Setup");
INSERT INTO missions(name) VALUES ("Slingshot");
INSERT INTO missions(name) VALUES ("Save Algie");
INSERT INTO missions(name) VALUES ("A Little Help");
INSERT INTO missions(name) VALUES ("Defend Bucky");
INSERT INTO missions(name) VALUES ("That Bitch");
INSERT INTO missions(name) VALUES ("The Candidate");
INSERT INTO missions(name) VALUES ("Halloween");
INSERT INTO missions(name) VALUES ("The Big Prank");
INSERT INTO missions(name) VALUES ("Help Gary");
INSERT INTO missions(name) VALUES ("Russell in the Hole");
INSERT INTO missions(name) VALUES ("The Diary");
INSERT INTO missions(name) VALUES ("Character Sheets");
INSERT INTO missions(name) VALUES ("Last Mintute Shopping");
INSERT INTO missions(name) VALUES ("Hattrick vs. Galloway");
INSERT INTO missions(name) VALUES ("Movie Tickets");
INSERT INTO missions(name) VALUES ("Carnival Date");
INSERT INTO missions(name) VALUES ("Prep Challenge");
INSERT INTO missions(name) VALUES ("The Eggs");
INSERT INTO missions(name) VALUES ("Race The Vale");
INSERT INTO missions(name) VALUES ("Beach Rumble");
INSERT INTO missions(name) VALUES ("Tad's House");
INSERT INTO missions(name) VALUES ("Boxing Challenge");
INSERT INTO missions(name) VALUES ("Dishonorable Fight");
INSERT INTO missions(name) VALUES ("Panty Raid");
INSERT INTO missions(name) VALUES ("Weed Killer");
INSERT INTO missions(name) VALUES ("Small Offences");
INSERT INTO missions(name) VALUES ("Christmas is Here");
INSERT INTO missions(name) VALUES ("Jealous Johnny");
INSERT INTO missions(name) VALUES ("Bait");
INSERT INTO missions(name) VALUES ("Tagging");
INSERT INTO missions(name) VALUES ("Wrong Part of Town");
INSERT INTO missions(name) VALUES ("Lola's Race");
INSERT INTO missions(name) VALUES ("The Tenements");
INSERT INTO missions(name) VALUES ("The Rumble");
INSERT INTO missions(name) VALUES ("Fighting Johnny Vincent");
INSERT INTO missions(name) VALUES ("Balls of Snow");
INSERT INTO missions(name) VALUES ("Miracle on Bullworth St.");
INSERT INTO missions(name) VALUES ("Nutcrackin");
INSERT INTO missions(name) VALUES ("Rudy the red Nosed Santa");
INSERT INTO missions(name) VALUES ("Cook's Crush");
INSERT INTO missions(name) VALUES ("Cook's Date");
INSERT INTO missions(name) VALUES ("Nerd Challenge");
INSERT INTO missions(name) VALUES ("Greaser Challenge");
INSERT INTO missions(name) VALUES ("Comic Klepto");
INSERT INTO missions(name) VALUES ("Glass House");
INSERT INTO missions(name) VALUES ("Discreet Deliveries");
INSERT INTO missions(name) VALUES ("Stronghold Assault");
INSERT INTO missions(name) VALUES ("Nerd Boss Fight");
INSERT INTO missions(name) VALUES ("Funhouse Fun");
INSERT INTO missions(name) VALUES ("Paparazzi");
INSERT INTO missions(name) VALUES ("Defender of the Castle");
INSERT INTO missions(name) VALUES ("Discretion Assured");
INSERT INTO missions(name) VALUES ("Nice Outfit");
INSERT INTO missions(name) VALUES ("The Big Game");
INSERT INTO missions(name) VALUES ("Jock Boss Fight");
INSERT INTO missions(name) VALUES ("Jock Challenge");
INSERT INTO missions(name) VALUES ("Here's to you Ms. Philips");
INSERT INTO missions(name) VALUES ("Galloway Away");
INSERT INTO missions(name) VALUES ("Making a Mark");
INSERT INTO missions(name) VALUES ("Rat's in the Library");
INSERT INTO missions(name) VALUES ("The Gym is Burning");
INSERT INTO missions(name) VALUES ("Finding Johnny Vincent");
INSERT INTO missions(name) VALUES ("Revenge on Mr. Burton");
INSERT INTO missions(name) VALUES ("Preppies Vandalized");
INSERT INTO missions(name) VALUES ("Go See The Principal");
INSERT INTO missions(name) VALUES ("Busting In, Part II");
INSERT INTO missions(name) VALUES ("Showdown at the Plant");
INSERT INTO missions(name) VALUES ("Complete Mayhem");
INSERT INTO missions(name) VALUES ("Final Showdown");
INSERT INTO missions(name) VALUES ("Cheating Time");
INSERT INTO missions(name) VALUES ("Smash It Up");
INSERT INTO missions(name) VALUES ("Townie Challenge");
INSERT INTO missions(name) VALUES ("Mailbox Armageddon");
INSERT INTO missions(name) VALUES ("The Collector");

INSERT INTO locations(name) VALUES ("Bullworth Academy");
INSERT INTO locations(name) VALUES ("Bullworth Town");
INSERT INTO locations(name) VALUES ("Old Bullworth Vale");
INSERT INTO locations(name) VALUES ("New Coventry");
INSERT INTO locations(name) VALUES ("Blue Skies Industrial Park");

INSERT INTO cliques(name) VALUES ("Non-Clique");
INSERT INTO cliques(name) VALUES ("Nerds");
INSERT INTO cliques(name) VALUES ("Bullies");
INSERT INTO cliques(name) VALUES ("Preps");
INSERT INTO cliques(name) VALUES ("Greasers");
INSERT INTO cliques(name) VALUES ("Jocks");
INSERT INTO cliques(name) VALUES ("Townies");

INSERT INTO students(name, clique_id) VALUES ("Jimmy Hopkins", 1);
INSERT INTO students(name, clique_id) VALUES ("Gary Smith", 1);
INSERT INTO students(name, clique_id) VALUES ("Pete Kowalski", 1);
INSERT INTO students(name, clique_id) VALUES ("Angie Ng", 1);
INSERT INTO students(name, clique_id) VALUES ("Christy Martin", 1);
INSERT INTO students(name, clique_id) VALUES ("Constantinos Brakus", 1);
INSERT INTO students(name, clique_id) VALUES ("Eunice Pound", 1);
INSERT INTO students(name, clique_id) VALUES ("Ivan Alexander", 1);
INSERT INTO students(name, clique_id) VALUES ("Gloria Jackson", 1);
INSERT INTO students(name, clique_id) VALUES ("Gordon Wakefield", 1);
INSERT INTO students(name, clique_id) VALUES ("Karen Johnson", 1);
INSERT INTO students(name, clique_id) VALUES ("Lance Jackson", 1);
INSERT INTO students(name, clique_id) VALUES ("Melody Adams", 1);
INSERT INTO students(name, clique_id) VALUES ("Pedro De La Hoya", 1);
INSERT INTO students(name, clique_id) VALUES ("Ray Hughes", 1);
INSERT INTO students(name, clique_id) VALUES ("Sheldon Thompson", 1);
INSERT INTO students(name, clique_id) VALUES ("Trevor Moore", 1);

INSERT INTO students(name, clique_id) VALUES ("Earnest Jones", 2);
INSERT INTO students(name, clique_id) VALUES ("Algernon Papadopoulos", 2);
INSERT INTO students(name, clique_id) VALUES ("Beatrice Trudeau", 2);
INSERT INTO students(name, clique_id) VALUES ("Bucky Pasteur", 2);
INSERT INTO students(name, clique_id) VALUES ("Cornelius Johnson", 2);
INSERT INTO students(name, clique_id) VALUES ("Donald Anderson", 2);
INSERT INTO students(name, clique_id) VALUES ("Fatty Johnson", 2);
INSERT INTO students(name, clique_id) VALUES ("Melvin O'Connor", 2);
INSERT INTO students(name, clique_id) VALUES ("Thad Carlson", 2);

INSERT INTO students(name, clique_id) VALUES ("Russell Northrop", 3);
INSERT INTO students(name, clique_id) VALUES ("Davis White", 3);
INSERT INTO students(name, clique_id) VALUES ("Ethan Robinson", 3);
INSERT INTO students(name, clique_id) VALUES ("Tom Gurney", 3);
INSERT INTO students(name, clique_id) VALUES ("Trent Northwick", 3);
INSERT INTO students(name, clique_id) VALUES ("Troy Miller", 3);
INSERT INTO students(name, clique_id) VALUES ("Wade Martin", 3);

INSERT INTO students(name, clique_id) VALUES ("Derby Harrington", 4);
INSERT INTO students(name, clique_id) VALUES ("Bif Taylor", 4);
INSERT INTO students(name, clique_id) VALUES ("Bryce Montrose", 4);
INSERT INTO students(name, clique_id) VALUES ("Chad Morris", 4);
INSERT INTO students(name, clique_id) VALUES ("Gord Vendome", 4);
INSERT INTO students(name, clique_id) VALUES ("Justin Vandervelde", 4);
INSERT INTO students(name, clique_id) VALUES ("Parker Ogilvie", 4);
INSERT INTO students(name, clique_id) VALUES ("Pinky Gauthier", 4);
INSERT INTO students(name, clique_id) VALUES ("Tad Spencer", 4);

INSERT INTO students(name, clique_id) VALUES ("Johnny Vincent", 5);
INSERT INTO students(name, clique_id) VALUES ("Hal Esposito", 5);
INSERT INTO students(name, clique_id) VALUES ("Lefty Mancini", 5);
INSERT INTO students(name, clique_id) VALUES ("Lola Lombardi", 5);
INSERT INTO students(name, clique_id) VALUES ("Lucky De Luca", 5);
INSERT INTO students(name, clique_id) VALUES ("Norton Williams", 5);
INSERT INTO students(name, clique_id) VALUES ("Peanut Romano", 5);
INSERT INTO students(name, clique_id) VALUES ("Ricky Pucino", 5);
INSERT INTO students(name, clique_id) VALUES ("Vance Medici", 5);

INSERT INTO students(name, clique_id) VALUES ("Ted Thompson", 6);
INSERT INTO students(name, clique_id) VALUES ("Bo Jackson", 6);
INSERT INTO students(name, clique_id) VALUES ("Casey Harris", 6);
INSERT INTO students(name, clique_id) VALUES ("Damon West", 6);
INSERT INTO students(name, clique_id) VALUES ("Dan Wilson", 6);
INSERT INTO students(name, clique_id) VALUES ("Juri Karamazov", 6);
INSERT INTO students(name, clique_id) VALUES ("Kirby Olsen", 6);
INSERT INTO students(name, clique_id) VALUES ("Luis Luna", 6);
INSERT INTO students(name, clique_id) VALUES ("Mandy Wiles", 6);

INSERT INTO students(name, clique_id) VALUES ("Edgar Munson", 7);
INSERT INTO students(name, clique_id) VALUES ("Clint", 7);
INSERT INTO students(name, clique_id) VALUES ("Duncan", 7);
INSERT INTO students(name, clique_id) VALUES ("Gurney", 7);
INSERT INTO students(name, clique_id) VALUES ("Jerry", 7);
INSERT INTO students(name, clique_id) VALUES ("Leon", 7);
INSERT INTO students(name, clique_id) VALUES ("Omar Romero", 7);
INSERT INTO students(name, clique_id) VALUES ("Otto Tyler", 7);
INSERT INTO students(name, clique_id) VALUES ("Zoe Taylor", 7);

INSERT INTO classes(name) VALUES ("Chemistry");
INSERT INTO classes(name) VALUES ("English");
INSERT INTO classes(name) VALUES ("Art");
INSERT INTO classes(name) VALUES ("Gym");
INSERT INTO classes(name) VALUES ("Biology");
INSERT INTO classes(name) VALUES ("Music");
INSERT INTO classes(name) VALUES ("Shop");
INSERT INTO classes(name) VALUES ("Photography");
INSERT INTO classes(name) VALUES ("Geography");
INSERT INTO classes(name) VALUES ("Math");
