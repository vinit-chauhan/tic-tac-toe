import axios from "axios";
import { useEffect, useState } from "react";
import { Text, View } from "react-native";
import { API_URL } from "../context/AuthContext";

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

  useEffect(() => {
    axios
      .get(`${API_URL}/profile`)
      .then((p) => setProfile(p.data))
      .catch((e) => console.log(e));
  }, []);

  return (
    <View>
      <Text>Login</Text>
      <Text>{profile.Username}</Text>
      <Text>{profile.ID}</Text>
    </View>
  );
};

export default Profile;
