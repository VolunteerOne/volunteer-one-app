import React, { useState, useEffect } from "react";
import { Pressable, StyleSheet, Alert } from "react-native";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";
const LikeButton = () => {
  const [liked, setLiked] = useState(false);

  useEffect(() => {
    if (liked) {
      Alert.alert("Event bookmarked");
    }
  }, [liked]);
  return (
    <Pressable
      onPress={() => {
        setLiked((isLiked) => !isLiked);
      }}
      style={styles.likeButton}
    >
      <MaterialCommunityIcons
        name={liked ? "heart" : "heart-outline"}
        size={32}
        color={liked ? "red" : "black"}
      />
    </Pressable>
  );
};

const styles = StyleSheet.create({
  likeButton: {
    marginLeft: "auto",
  },
});

export default LikeButton;
