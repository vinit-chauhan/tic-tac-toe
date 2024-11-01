import { useState } from "react";
import {
  SafeAreaView,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from "react-native";
import { useAuth } from "../context/AuthContext";
import { NativeStackNavigationProp } from "@react-navigation/native-stack";
import { RootStackParamList } from "../navigarionTypes";
import { style } from "./styles/root";

type SignupScreenNavigationProp = NativeStackNavigationProp<
  RootStackParamList,
  "Signup"
>;

type Props = {
  navigation: SignupScreenNavigationProp;
};

const Signup = ({ navigation }: Props) => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [email, setEmail] = useState("");
  const [error, setError] = useState("");

  const { onRegister } = useAuth();

  const handleRegistration = async () => {
    if (!username || !password) {
      return;
    }

    const result = await onRegister!(username, password, email);
    if (result && result.error) {
      alert(result.msg);
    }
  };

  const validateUsername = (username: string) => {
    if (username.length < 3) {
      setError("Username must be at least 3 characters long");
    } else {
      setError("");
    }
  };
  const validatePassword = (password: string) => {
    if (password.length < 4) {
      setError("Password must be at least 8 characters long");
    } else {
      setError("");
    }
  };

  return (
    <SafeAreaView style={[style.container, style.backdrop]}>
      <View style={[style.container, style.formGroup]}>
        <TextInput
          style={style.textBox}
          placeholder="Username"
          autoComplete="username"
          onChangeText={(username) => setUsername(username)}
          onBlur={(e) => {
            validateUsername(e.nativeEvent.text);
          }}
        />

        <TextInput
          style={style.textBox}
          id="email"
          placeholder="Email"
          autoComplete="email"
          onChangeText={(email) => setEmail(email)}
        />

        <TextInput
          style={style.textBox}
          id="password"
          placeholder="Password"
          autoComplete="current-password"
          secureTextEntry={true}
          onChangeText={(password) => setPassword(password)}
          onBlur={(e) => {
            validatePassword(e.nativeEvent.text);
          }}
        />

        {error && <Text style={{ color: "red" }}>{error}</Text>}

        <TouchableOpacity
          onPress={handleRegistration}
          disabled={error.length > 0}
          style={[style.loginBtn, { opacity: error.length > 0 ? 0.5 : 1 }]}
        >
          <Text style={style.text}>Sign Up</Text>
        </TouchableOpacity>

        <TouchableOpacity onPress={() => navigation.navigate("Login")}>
          <Text style={style.link}>Already a member?</Text>
        </TouchableOpacity>
      </View>
    </SafeAreaView>
  );
};

export default Signup;
