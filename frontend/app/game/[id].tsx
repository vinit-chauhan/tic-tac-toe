import axios from "axios";
import { useEffect, useState } from "react";
import { View, Text, TouchableOpacity } from "react-native";
import { style } from "../screens/styles/root";
import { API_URL } from "../context/AuthContext";
import { useLocalSearchParams } from "expo-router";

const GameScreen = () => {
  const [game, setGame] = useState(Array(9).fill(null));
  const [winner, setWinner] = useState<string | null>(null);
  const [waiting, setWaiting] = useState<boolean>(true);

  const { id } = useLocalSearchParams();
  const gameId = Array.isArray(id) ? id[0] : id;

  const updateGameState = (id: string) => {
    axios
      .get(`${API_URL}/game/${id}`)
      .then((resp) => {
        setGame(resp.data["board"]);
        if (resp.data["winner"]) {
          setWinner(resp.data["winner"]);
        }
        !resp.data["waiting"] && setWaiting(false)
      })
      .catch((error) => console.log(error));
  };

  useEffect(() => {
    // TODO: Use long polling
    const interval = setInterval(() => {
      updateGameState(gameId);
    }, 10000);

    updateGameState(gameId);
    return () => clearInterval(interval);
  }, [winner]);

  return (
    <View style={style.container}>
      {waiting ? (
        <View>
          <Text>Waiting for another player</Text>
          <Text>Game ID: {gameId}</Text>
        </View>
      ) : (
        <View>
          <Text>Game</Text>
          {winner ? (
            <Text>{winner} Won!!</Text>
          ) : (
            <View style={{ flexDirection: "column" }}>
              {[0, 1, 2].map((row) => (
                <View
                  key={row}
                  style={{ flexDirection: "row", backgroundColor: "black" }}
                >
                  {[0, 1, 2].map((col) => (
                    <TouchableOpacity
                      onPress={() => {
                        axios
                          .put(`${API_URL}/game/${gameId}/move`, {
                            row: row,
                            column: col,
                          })
                          .then((resp) => {
                            console.log(resp.data);

                            setGame(resp.data["board"]);

                            if (resp.data["winner"]) {
                              console.log(`Winner is: ${resp.data["winner"]}`);
                              setWinner(resp.data["winner"]);
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
      )}
    </View>
  );
};

export default GameScreen;
