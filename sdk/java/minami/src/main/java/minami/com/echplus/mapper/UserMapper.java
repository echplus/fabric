package minami.com.echplus.mapper;

import java.util.List;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Select;

import minami.com.echplus.domain.User;

@Mapper
public interface UserMapper {

    @Select("SELECT row, id, enrollmentId FROM Users WHERE row <= #{row}")
    List<User> getUsers0(@Param("row") int row);

    List<User> getUsers1(Integer row);

}
