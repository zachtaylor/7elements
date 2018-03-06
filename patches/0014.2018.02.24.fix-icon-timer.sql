-- fix refs to /img/icon/timer.png
UPDATE cards_powers_texts SET description="<b><img class='se-symbol' src='/img/icon/element-3.png'>+<img class='se-symbol' src='/img/icon/timer.20px.png'>:</b>'boen' gets +1 <img src='/img/icon/attack.20px.png'> and +1 <img src='/img/icon/life.20px.png'>" WHERE cardid=3;
UPDATE cards_powers_texts SET description="<b><img class='se-symbol' src='/img/icon/timer.20px.png'>:</b> Create a clone of target body, then you gain 1 life" WHERE cardid=48;
