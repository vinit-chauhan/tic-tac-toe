import { useState } from "react";
import {
  SafeAreaView,
  StyleSheet,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from "react-native";
import { useAuth } from "../context/AuthContext";

import { NativeStackNavigationProp } from "@react-navigation/native-stack";
import { RootStackParamList } from "../navigarionTypes";
import { style } from "./styles/root";

type LoginScreenNavigationProp = NativeStackNavigationProp<
  RootStackParamList,
  "Login"
>;

type Props = {
  navigation: LoginScreenNavigationProp;
};

const Login = ({ navigation }: Props) => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const { onLogin } = useAuth();

  const validateUsername = (username: string) => {
    if (username === "") {
      setError("Username is required");
    } else {
      setError("");
    }
  };

  const validatePassword = (password: string) => {
    if (password === "") {
      setError("Password is required");
    } else {
      setError("");
    }
  };

  const handleLogin = async () => {
    const result = await onLogin!(username, password);
    if (result && result.error) {
      alert(result.msg);
    }
  };

  return (
    <SafeAreaView style={[style.container, style.backdrop]}>
      <View style={[style.container, style.formGroup]}>
        <TextInput
          style={style.textBox}
          placeholder="Username"
          onBlur={(e) => validateUsername(e.nativeEvent.text)}
          onChangeText={(username) => setUsername(username)}
        />

        <TextInput
          style={style.textBox}
          placeholder="Password"
          autoComplete="current-password"
          secureTextEntry={true}
          onBlur={(e) => validatePassword(e.nativeEvent.text)}
          onChangeText={(password) => setPassword(password)}
        />

        {error && <Text style={{ color: "red" }}>{error}</Text>}

        <TouchableOpacity
          onPress={handleLogin}
          disabled={error.length > 0}
          style={[style.loginBtn, { opacity: error.length > 0 ? 0.5 : 1 }]}
        >
          <Text style={style.text}>Sign in</Text>
        </TouchableOpacity>

        <TouchableOpacity onPress={() => navigation.navigate("Signup")}>
          <Text style={style.link}>Don't have an account?</Text>
        </TouchableOpacity>
      </View>
    </SafeAreaView>
  );
};

export default Login;
