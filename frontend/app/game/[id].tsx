import axios from "axios";
import { useEffect, useState } from "react";
import { View, Text, TouchableOpacity } from "react-native";
import { style } from "../screens/styles/root";
import { API_URL } from "../context/AuthContext";

const GameScreen = () => {
  const [game, setGame] = useState(Array(9).fill(null));
  const [winner, setWinner] = useState<string | null>(null);

  const updateGameState = (id: string) => {
    axios
      .get(`${API_URL}/game/${id}`)
      .then((resp) => {
        setGame(resp.data["board"]);
        if (resp.data["winner"]) {
          setWinner(resp.data["winner"]);
        }
      })
      .catch((error) => console.log(error));
  };

  useEffect(() => {
    const interval = setInterval(() => {
      if (!winner) {
        updateGameState("3735928558");
      }
    }, 1000);

    updateGameState("3735928558");
    return () => clearInterval(interval);
  }, [winner]);

  return (
    <View style={style.container}>
      <View>
        <Text>Game</Text>
      </View>
      {winner ? (
        <Text> {winner} Won!! </Text>
      ) : (
        <View style={{ flexDirection: "column" }}>
          <Text> {winner} Won!! </Text>
          {[0, 1, 2].map((row) => (
            <View
              key={row}
              style={{ flexDirection: "row", backgroundColor: "black" }}
            >
              {[0, 1, 2].map((col) => (
                <TouchableOpacity
                  onPress={() => {
                    axios
                      .put(`${API_URL}/game/3735928558/move`, {
                        row: row,
                        column: col,
                      })
                      .then((resp) => {
                        console.log(resp.data);

                        setGame(resp.data["board"]);

                        if (resp.data["winner"]) {
                          console.log(`Winner is: ${resp.data["winner"]}`);
                          setWinner(winner);
                        }
                      })
                      .catch((err) => console.log(`ERROR: ${err}`));
                  }}
                  key={col}
                  style={{
                    width: 100,
                    height: 100,
                    margin: 1,
                    justifyContent: "center",
                    alignItems: "center",
                    backgroundColor: "white",
                  }}
                >
                  <Text>
                    {game[row * 3 + col] == 0
                      ? ""
                      : game[row * 3 + col] == 1
                      ? "X"
                      : "O"}
                  </Text>
                </TouchableOpacity>
              ))}
            </View>
          ))}
        </View>
      )}
    </View>
  );
};

export default GameScreen;
