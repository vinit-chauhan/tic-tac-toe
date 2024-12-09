import axios from "axios";
import { useEffect, useState } from "react";
import { Text, TextInput, TouchableOpacity, View } from "react-native";
import { API_URL } from "../context/AuthContext";
import { style } from "./styles/root";
import { router } from "expo-router";

type UserProfile = {
  ID: number;
  Username: string;
  Email: string;
};

const Profile = () => {
  const [profile, setProfile] = useState<UserProfile>({
    ID: -1,
    Username: "",
    Email: "",
  });

  const [gameId, setGameId] = useState<string>("");
  const [error, setError] = useState<string>("");

  useEffect(() => {
    axios
      .get(`${API_URL}/profile`)
      .then((p) => setProfile(p.data))
      .then(() => {
        axios
          .post(`${API_URL}/game`, {
            player1: profile.ID,
            player2: 0,
          })
          .then((g) => setGameId(g.data.ID));
      })
      .catch((e) => console.log(e));
  }, []);

  return (
    <View style={style.container}>
      <Text>Login</Text>
      <Text>{profile.Username}</Text>

      <View style={[style.container, style.formGroup]}>
        <TextInput
          style={style.textBox}
          placeholder="Game ID"
          onChangeText={(id) => setGameId(id)}
        />
        <TouchableOpacity
          style={style.loginBtn}
          onPress={() =>
            axios
              .post(`${API_URL}/game/${gameId}/join`)
              .then(() =>
                router.push({
                  pathname: "/game/[id]",
                  params: { id: gameId },
                })
              )
              .catch((err) => {
                setError(`Unable to join game:${gameId} with`);
                console.log(err);
              })
          }
        >
          Join Game
        </TouchableOpacity>

        <TouchableOpacity
          style={style.link}
          onPress={() =>
            axios
              .post(`${API_URL}/game`)
              .then((resp) =>
                router.push({
                  pathname: "/game/[id]",
                  params: { id: resp.data["game_id"] },
                })
              )
              .catch((err) => console.log(err))
          }
        >
          <Text>Start New Game</Text>
        </TouchableOpacity>
      </View>
    </View>
  );
};

export default Profile;
